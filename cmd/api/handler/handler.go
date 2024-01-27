package handler

import (
	"encoding/json"
	"errors"
	"github.com/julienschmidt/httprouter"
	"net/http"

	"github/bthari/tinble/internal/constant"
	response "github/bthari/tinble/internal/dto/response"
	"github/bthari/tinble/internal/usecase"
	"github/bthari/tinble/pkg/config"
)

type HandlerInterface interface {
	Ping(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	RegisterUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	SignIn(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}

type Handler struct {
	UseCase usecase.UseCaseInterface
	Config  *config.Config
}

func NewHandler(config *config.Config, uc usecase.UseCase) Handler {
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
