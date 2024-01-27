package usecase

import (
	"context"

	"github/bthari/tinble/internal/dto/request"
	"github/bthari/tinble/internal/model"
	"github/bthari/tinble/internal/store"
)

type UseCaseInterface interface {
	InsertNewUser(ctx context.Context, user *model.User) (err error)
	AuthenticateUser(ctx context.Context, request *request.SignInRequest) (err error)
}

type UseCase struct {
	Store store.StoreInterface
}

func NewUseCase(store store.StoreInterface) UseCase {
	return UseCase{
		Store: store,
	}
}
