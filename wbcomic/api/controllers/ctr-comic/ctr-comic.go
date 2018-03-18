package ctr_comic

import (
	"wbcomic/api/client"
	"wbcomic/pb"
	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context)  {

	res,_ :=pb.NewComicServiceClient(client.NewRpcConn()).GetComicListByPage(c, &pb.ComicFilterRequest{Status:1})

	c.JSON(200,gin.H{
		"code":1,
		"message":"ok",
		"data":res,
	})
}
