package main

import (
	"fmt"
	handler "github/bthari/tinble/app/cmd/api/handler"
	router "github/bthari/tinble/app/cmd/api/router"
	"log"
	"net/http"
)

func main() {
	h := handler.NewHandler()
	r := router.Init(h)
	addr := ":8080"

	fmt.Println("web server running on: ", addr)
	err := http.ListenAndServe(addr, r)
	if err != nil {
		log.Fatalf("failed to start, err: %v", err)
	}
}
