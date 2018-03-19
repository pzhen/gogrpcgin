package main

import(
	"net"
	"errors"
	"runtime"
	"context"
	"net/http"
	"time"

	"gogrpcgin/conf"
	"gogrpcgin/utils"
	"gogrpcgin/service"

	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc"
	"golang.org/x/net/trace"
)

// dev/pro
const Env = "dev"

// rpc run
func main() {

	// init config
	conf.InitConfig(Env,"rpc")

	//new server engine
	server := newServer()

	// register service
	service.ServiceReg(server)

	// run
	serviceRun(server)
}

// service run
func serviceRun(server *grpc.Server) {

	address := conf.Conf.App.Rpc.RpcAddr
	lis, err := net.Listen("tcp", address)
	if err != nil {
		utils.LogFatalfError(err)
	}
	utils.LogPrint("RPC server listen on %s\n", address)

	// open trace
	startTrace()

	// 开启服务
	reflection.Register(server)
	if err := server.Serve(lis); err != nil {
		utils.LogFatalfError(err)
	}
}

// open trace
// like: http://localhost:50052/debug/requests
// like: http://localhost:50052/debug/events
func startTrace() {

	grpc.EnableTracing = true

	trace.AuthRequest = func(req *http.Request) (any, sensitive bool) {
		return true, true
	}

	var traceAddr = conf.Conf.App.Rpc.RpcTraceAddr

	utils.LogPrint("Trace listen on %s\n",traceAddr)
	go http.ListenAndServe(traceAddr, nil)
}


// new server engine
func newServer() *grpc.Server{

	return grpc.NewServer(grpc.UnaryInterceptor(composeUnaryServerInterceptors(serverRecoveryInterceptorHandle,serverDurationInterceptorHandle)))
}

// server recovery interceptor handle
func serverRecoveryInterceptorHandle(ctx context.Context,req interface{},info *grpc.UnaryServerInfo,handler grpc.UnaryHandler)(
	interface{}, error) {

	defer func() (err error){
		if r := recover(); r != nil {
			str, ok := r.(string)

			if ok {
				err = errors.New(str)
			} else {
				err = errors.New("panic!!!")
			}

			//打印堆栈
			stack := make([]byte, 1024*8)
			stack = stack[:runtime.Stack(stack, false)]

			utils.LogPrint("[Recovery] panic recovered:\n\n%s\n\n%s" , r,stack)
		}

		return
	}()

	return handler(ctx, req)
}

// server duration interceptor handle
func serverDurationInterceptorHandle(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	start := time.Now()
	resp, err := handler(ctx, req)

	utils.LogPrint("invoke server method=%s duration=%s error=%v", info.FullMethod, time.Since(start), err)

	return resp, err
}

// multi interceptors
func composeUnaryServerInterceptors(interceptor, next grpc.UnaryServerInterceptor) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (interface{}, error) {

		return interceptor(ctx, req, info,
			func(nextCtx context.Context, nextReq interface{}) (interface{}, error) {

				return next(nextCtx, nextReq, info, handler)
			})
	}
}

