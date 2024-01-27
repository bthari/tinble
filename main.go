package main

import (
	"fmt"
	"log"
	"net/http"

	"github/bthari/tinble/cmd/api/handler"
	"github/bthari/tinble/cmd/api/router"
	"github/bthari/tinble/internal/store"
	"github/bthari/tinble/internal/usecase"
	config "github/bthari/tinble/pkg/config"
)

func main() {
	config := config.GetConfig()
	s := store.NewStore(store.InitMongo(config.Mongo))
	uc := usecase.NewUseCase(&s)
	h := handler.NewHandler(config, uc)
	r := router.Init(h)
	addr := ":8080"

	fmt.Println("web server running on: ", addr)
	err := http.ListenAndServe(addr, r)
	if err != nil {
		log.Fatalf("failed to start, err: %v", err)
	}
}
