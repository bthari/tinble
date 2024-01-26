package handler

import (
	"encoding/json"
	"errors"
	"github.com/julienschmidt/httprouter"
	"github/bthari/tinble/app/internal/constant"
	response "github/bthari/tinble/app/internal/dto/response"
	"github/bthari/tinble/app/internal/store"
	"github/bthari/tinble/app/internal/usecase"
	"github/bthari/tinble/app/pkg/config"
	"net/http"
)

type HandlerInterface interface {
	Ping(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	RegisterUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	SignIn(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}

type Handler struct {
	UseCase usecase.UseCase
	Config  *config.Config
}

func NewHandler(config *config.Config) Handler {
	s := store.NewStore(store.InitStore(config.Mongo))
	uc := usecase.NewUseCase(&s)

	return Handler{
		UseCase: uc,
		Config:  config,
	}
}

func WriteSuccessResponse(w http.ResponseWriter, data interface{}) {
	resp := &response.Response{
		Msg:  "success",
		Data: data,
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		json.NewEncoder(w).Encode(&response.Response{Msg: "failed to encode response"})
	}
}

func WriteErrorResponse(w http.ResponseWriter, err error) {
	statusCode, errorCode := GetErrorResponse(err)
	resp := &response.Response{
		Code: errorCode,
		Msg:  err.Error(),
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		json.NewEncoder(w).Encode(
			&response.Response{
				Code: errorCode,
				Msg:  err.Error(),
			})
	}
}

func GetErrorResponse(err error) (statusCode, errorCode int) {
	statusCode = http.StatusInternalServerError
	errorCode = 500001

	switch {
	case errors.Is(err, constant.ErrEmailExist):
		statusCode = http.StatusBadRequest
		errorCode = 400001
	case errors.Is(err, constant.ErrUsernameExist):
		statusCode = http.StatusBadRequest
		errorCode = 400002
	case errors.Is(err, constant.ErrUserNotExist):
		statusCode = http.StatusBadRequest
		errorCode = 400003
	case errors.Is(err, constant.ErrPasswordNotMatch):
		statusCode = http.StatusBadRequest
		errorCode = 400004
	}

	return
}
