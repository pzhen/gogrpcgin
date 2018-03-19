package routers

import(
	"github.com/gin-gonic/gin"
	"gogrpcgin/conf"
	"gogrpcgin/api/controllers/ctr-comic"
	"time"
	jwt_lib "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/contrib/jwt"
	"github.com/ekyoung/gin-nice-recovery"
	"github.com/DeanThompson/ginpprof"
)

func InitRun()  {

	engine := gin.Default()
	//http://127.0.0.1:8080/debug/pprof/
	ginpprof.Wrap(engine)

	engine.Use(nice.Recovery(recoveryHandler))

	// 公开路由
	public := engine.Group("/rpc")
	{
		public.POST("/auth", createTokenHandle)
		public.GET("/comic/test",ctr_comic.Test)
	}

	// 非公开路由
	private := engine.Group("/rpc/private")
	{
		private.Use(jwt.Auth(conf.Conf.App.Api.ApiSecretKey))

		private.POST("/comic/show", func(c *gin.Context) {
			//library.ProcessRPC(c, &c_comic.ComicRPC{})
		})

		private.GET("/test",func(c *gin.Context) {
			panic("sssssssss")
			c.JSON(200, gin.H{"message": "Hello from private"})
		})
	}




	// service run
	if conf.Conf.App.Api.ApiTls == true {
		engine.RunTLS(conf.Conf.App.Api.ApiTlsAddr,"keys/server.crt","keys/server.key")
	}else {
		engine.Run(conf.Conf.App.Api.ApiAddr)
	}
}

// 生成 token
func createTokenHandle(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")

	if username == conf.Conf.App.Api.ApiUsername && password == conf.Conf.App.Api.ApiPassword {
		// Create the token
		token := jwt_lib.New(jwt_lib.GetSigningMethod("HS256"))
		// Set some claims
		token.Claims = jwt_lib.MapClaims{
			"Id":  "Christopher",
			"exp": time.Now().Add(time.Hour * 1).Unix(),
		}
		// Sign and get the complete encoded token as a string
		tokenString, err := token.SignedString([]byte(conf.Conf.App.Api.ApiSecretKey))
		if err != nil {
			c.JSON(500, gin.H{"code":0, "message": "Could not generate token"})
			return
		}
		c.JSON(200, gin.H{"code":0,"message":"ok","data":gin.H{"token":tokenString}})
		return
	}else {

		c.JSON(200, gin.H{"code":0, "message": "username or password is wrong"})
		return
	}


}

// 捕捉 panic
func recoveryHandler(c *gin.Context, err interface{}) {
	c.JSON(500, gin.H{"code":0, "message": err})
}