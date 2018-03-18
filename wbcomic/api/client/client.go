package client

import(
	"sync"
	"google.golang.org/grpc"
	"wbcomic/conf"
)

func NewRpcConn() *grpc.ClientConn{

	var address = conf.Conf.App.Rpc.RpcAddr

	p := &sync.Pool{
		New: func() interface {}{
			conn,_ := grpc.Dial(address, grpc.WithInsecure())
			return  conn
		},
	}

	return p.Get().(*grpc.ClientConn)
}
