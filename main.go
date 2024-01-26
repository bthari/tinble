package main

import (
	"fmt"
	"github/bthari/tinble/app/cmd/api/handler"
	"github/bthari/tinble/app/cmd/api/router"
	config "github/bthari/tinble/app/pkg/config"
	"log"
	"net/http"
)

func main() {
	config := config.GetConfig()
	h := handler.NewHandler(config)
	r := router.Init(h)
	addr := ":8080"

	fmt.Println("web server running on: ", addr)
	err := http.ListenAndServe(addr, r)
	if err != nil {
		log.Fatalf("failed to start, err: %v", err)
	}
}
