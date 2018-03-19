package main

import(
	"wbcomic/conf"
	"wbcomic/api/routers"
)

// dev/pro
const Env = "dev"

func main()  {
	conf.InitConfig(Env,"api")
	routers.InitRun()
}

