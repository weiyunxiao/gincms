package main

import (
	"gincms/app"
	"gincms/app/bootstrap"
	"gincms/app/http"
	"gincms/app/http/router"
	"log"
)

func main() {
	//初始化应用组件，加载配置
	bootstrap.AppInit()
	//创建gin引擎
	r := http.CreateGinServer()
	//路由定义
	router.InitRoute(r)

	port := app.Config.Http.AdminApiPort
	log.Println(app.Config.App.Name + "运行在端口:" + port)
	if err := r.Run(port); err != nil {
		log.Fatalln(err)
	}
}
