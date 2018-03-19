package main

import(
	"gogrpcgin/conf"
	"gogrpcgin/api/routers"
)

// dev/pro
const Env = "dev"

func main()  {
	conf.InitConfig(Env,"api")
	routers.InitRun()
}

