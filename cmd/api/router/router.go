package router

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github/bthari/tinble/cmd/api/handler"
)

func Init(handler handler.Handler) http.Handler {
	router := httprouter.New()

	router.GET("/ping", handler.Ping)
	router.POST("/register", handler.RegisterUser)
	router.POST("/sign-in", handler.SignIn)

	return router
}
