package usecase

import (
	"context"
	"github/bthari/tinble/app/internal/constant"
	"github/bthari/tinble/app/internal/dto/request"
	"github/bthari/tinble/app/internal/model"
	"github/bthari/tinble/app/internal/util"
	"go.mongodb.org/mongo-driver/mongo"
)

func (uc UseCase) InsertNewUser(ctx context.Context, newUser *model.User) (err error) {
	user, err := uc.Store.FindUserByUsernameOrEmail(ctx, newUser.Username, newUser.Email)
	if err != nil && err != mongo.ErrNoDocuments {
		return err
	}

	if user != nil {
		if user.Username == newUser.Username {
			return constant.ErrUsernameExist
		}

		if user.Email == newUser.Email {
			return constant.ErrEmailExist
		}
	}

	err = uc.Store.InsertOne(ctx, newUser)
	if err != nil {
		return err
	}

	return
}

func (uc UseCase) AuthenticateUser(ctx context.Context, request *request.SignInRequest) (err error) {
	user, err := uc.Store.FindUserByUsernameOrEmail(ctx, request.Username, request.Email)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return constant.ErrUserNotExist
		}
		return err
	}

	valid := util.ValidatePassword(user.Password, request.Password)
	if !valid {
		return constant.ErrPasswordNotMatch
	}

	return nil
}
