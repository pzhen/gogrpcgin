package service

import (
	"google.golang.org/grpc"
	"gogrpcgin/pb"
	"gogrpcgin/service/srv-comic"
)

func ServiceReg(server *grpc.Server)  {

	pb.RegisterComicServiceServer(server,&srv_comic.ComicInfo{})

}
