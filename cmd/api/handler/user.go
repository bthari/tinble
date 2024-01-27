package handler

import (
	"encoding/json"
	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/julienschmidt/httprouter"
	"github/bthari/tinble/internal/dto/request"
	"github/bthari/tinble/internal/dto/response"
	"github/bthari/tinble/internal/model"
	"github/bthari/tinble/internal/util"
	"net/http"
	"time"
)

func (h *Handler) RegisterUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	body := &request.RegisterUserBody{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&body)
	if err != nil {
		WriteErrorResponse(w, err)
		return
	}

	hashedPass, err := util.HashPassword(body.Password)
	if err != nil {
		WriteErrorResponse(w, err)
		return
	}

	newUser := &model.User{
		Username: body.Username,
		Email:    body.Email,
		Password: hashedPass,
	}

	err = h.UseCase.InsertNewUser(r.Context(), newUser)
	if err != nil {
		WriteErrorResponse(w, err)
		return
	}

	WriteSuccessResponse(w, nil)
	return
}

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	body := &request.SignInRequest{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&body)
	if err != nil {
		WriteErrorResponse(w, err)
		return
	}

	err = h.UseCase.AuthenticateUser(r.Context(), body)
	if err != nil {
		WriteErrorResponse(w, err)
		return
	}

	token, err := h.generateJWT(body)
	if err != nil {
		WriteErrorResponse(w, err)
		return
	}

	WriteSuccessResponse(w, &response.JWTTokenResponse{
		Token: token,
	})

	return
}

func (h *Handler) generateJWT(body *request.SignInRequest) (token string, err error) {
	expirationTime := time.Now().Add(15 * time.Minute)
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": body.Username,
			"email":    body.Email,
			"exp":      expirationTime,
		})

	token, err = jwtToken.SignedString([]byte(h.Config.Auth.JWTSecret))
	if err != nil {
		return "", err
	}

	return token, err
}
