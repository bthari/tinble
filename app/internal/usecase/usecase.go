package usecase

import (
	"context"
	"github/bthari/tinble/app/internal/model"
	"github/bthari/tinble/app/internal/store"
)

type UseCaseInterface interface {
	InsertNewUser(ctx context.Context, user *model.User) (err error)
	AuthenticateUser(ctx context.Context, username string) (user *model.User, err error)
}

type UseCase struct {
	Store store.StoreInterface
}

func NewUseCase(store store.StoreInterface) UseCase {
	return UseCase{
		Store: store,
	}
}
