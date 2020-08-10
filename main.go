package main

import (
	"fmt"
	"net/http"

	"gin-example/routers"
	"gin-example/util"
)

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", util.HTTPPort),
		Handler:        router,
		ReadTimeout:    util.ReadTimeout,
		WriteTimeout:   util.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
