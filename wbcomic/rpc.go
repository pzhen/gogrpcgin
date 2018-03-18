package main

import(
	"net"
	"errors"
	"runtime"

	"wbcomic/conf"
	"wbcomic/utils"

	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"wbcomic/service"
	"golang.org/x/net/trace"
	"net/http"
	"google.golang.org/grpc/grpclog"
)

// dev/pro
const Env = "dev"

// recover func handle
var recoverFunc grpc_recovery.RecoveryHandlerFunc = func(p interface{}) (err error) {

	str, ok := p.(string)
	if ok {
		err = errors.New(str)
	} else {
		err = errors.New("panic!!!")
	}

	//打印堆栈
	stack := make([]byte, 1024*8)
	stack = stack[:runtime.Stack(stack, false)]

	utils.LogPrint("[Recovery] panic recovered:\n\n%s\n\n%s" , p,stack)
	return err
}

// service run
func serviceRun(server *grpc.Server)  {
	// init config
	conf.InitConfig(Env,"rpc")

	lis, err := net.Listen("tcp", conf.Conf.App.Rpc.RpcAddr)
	if err != nil {
		utils.LogFatalfError(err)
	}

	startTrace()

	// 开启服务
	reflection.Register(server)
	if err := server.Serve(lis); err != nil {
		utils.LogFatalfError(err)
	}
}

// open trace
// http://localhost:50052/debug/requests
// http://localhost:50052/debug/events
func startTrace() {
	grpc.EnableTracing = true

	trace.AuthRequest = func(req *http.Request) (any, sensitive bool) {
		return true, true
	}
	go http.ListenAndServe(":50052", nil)
	grpclog.Infoln("Trace listen on 50051")
}


func main() {

	// interceptor
	opts := []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandler(recoverFunc),
	}
	server := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_recovery.UnaryServerInterceptor(opts...),
		),
		grpc_middleware.WithStreamServerChain(
			grpc_recovery.StreamServerInterceptor(opts...),
		),
	)

	// register service
	service.ServiceReg(server)

	// run
	serviceRun(server)
}
