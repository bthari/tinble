package handler

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (h *Handler) Ping(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	data, _ := json.Marshal(&struct {
		Status int
	}{
		Status: http.StatusOK,
	})

	if _, err := w.Write(data); err != nil {
		return
	}

	return
}
