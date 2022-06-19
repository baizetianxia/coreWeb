package main

import (
	"net/http"

	"github.com/baizetianxia/coreWeb/framework"
)

func main() {
	core := framework.NewCore()
	registerRouter(core)

	server := &http.Server{
		Addr:    ":8880",
		Handler: core,
	}

	server.ListenAndServe()
}
