package router

import (
	"github.com/julienschmidt/httprouter"
	"github/bthari/tinble/app/cmd/api/handler"
	"net/http"
)

func Init(handler handler.Handler) http.Handler {
	router := httprouter.New()

	router.GET("/ping", handler.Ping)
	router.POST("/register", handler.RegisterUser)
	router.POST("/sign-in", handler.SignIn)

	return router
}
