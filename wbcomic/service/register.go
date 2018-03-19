package service

import (
	"google.golang.org/grpc"
	"wbcomic/pb"
	"wbcomic/service/srv-comic"
)

func ServiceReg(server *grpc.Server)  {

	pb.RegisterComicServiceServer(server,&srv_comic.ComicInfo{})

}
