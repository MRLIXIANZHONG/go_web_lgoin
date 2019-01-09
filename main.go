package main

import (
	"fmt"
	"net/http"

	"web_test/pkg/setting"
	"web_test/router"
)

func main() {

	router := routers.InitRouter()//初始化路由

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()


}
