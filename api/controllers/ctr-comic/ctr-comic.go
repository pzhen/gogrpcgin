package ctr_comic

import (
	"gogrpcgin/api/client"
	"gogrpcgin/pb"
	"github.com/gin-gonic/gin"
	"gogrpcgin/core"
)


func Test(c *gin.Context)  {

	res,_ :=pb.NewComicServiceClient(client.NewRpcConn()).GetComicListByPage(c, &pb.ComicFilterRequest{Status:1})

	var user core.EsResponse
	core.MasterES("author","author").Query(`{"query": {"bools": {"should": {"regexp": {"sina_nickname": ".*果.*"}}}}}`,&user)

	c.JSON(200,gin.H{
		"code":1,
		"message":"ok",
		"data":res,
		"data2":user,
	})
}
