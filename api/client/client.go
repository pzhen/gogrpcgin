package client

import(
	"sync"
	"google.golang.org/grpc"
	"gogrpcgin/conf"
	"time"
	"context"
	"gogrpcgin/utils"
)


func NewRpcConn() *grpc.ClientConn{

	var address = conf.Conf.App.Rpc.RpcAddr

	p := &sync.Pool{
		New: func() interface {}{
			conn,_ := grpc.Dial(address, grpc.WithInsecure(),WithClientDurativeInterceptor())
			return  conn
		},
	}

	return p.Get().(*grpc.ClientConn)
}

func WithClientDurativeInterceptor() grpc.DialOption {

	return grpc.WithUnaryInterceptor(clientDurativeInterceptorHandle)
}

func clientDurativeInterceptorHandle(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {

	start := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)
	utils.LogPrint("invoke remote method=%s duration=%s error=%v", method, time.Since(start), err)
	return err
}
