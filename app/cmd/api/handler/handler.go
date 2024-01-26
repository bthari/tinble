package handler

import (
	"encoding/json"
	"github/bthari/tinble/app/internal/dto/response"
	"github/bthari/tinble/app/pkg/config"
	"net/http"
)

type Handler struct {
}

func NewHandler(config *config.Config) Handler {
	return Handler{}
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

	return
}
