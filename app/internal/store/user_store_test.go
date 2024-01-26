package store_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github/bthari/tinble/app/internal/model"
	store "github/bthari/tinble/app/internal/store"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"testing"
)

func TestFindOne(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	type args struct {
		expectedUser *model.User
	}
	testCases := []struct {
		name         string
		args         args
		expectations func(mt *mtest.T, args args) (res *model.User, err error)
		result       func(args args, res *model.User, err error)
	}{
		{
			name: "Match username",
			args: args{
				expectedUser: &model.User{
					ID:       primitive.NewObjectID(),
					Username: "user",
					Email:    "user@gmail.com",
					Password: "hashed",
				},
			},
			expectations: func(mt *mtest.T, args args) (res *model.User, err error) {
				s := store.NewStore(mt.DB)
				mt.AddMockResponses(mtest.CreateCursorResponse(1, "test.test", mtest.FirstBatch, bson.D{
					{"_id", args.expectedUser.ID},
					{"username", args.expectedUser.Username},
					{"email", args.expectedUser.Email},
					{"password", args.expectedUser.Password},
				}))

				return s.FindUserByUsernameOrEmail(context.TODO(), args.expectedUser.Username, "not email")
			},
			result: func(args args, res *model.User, err error) {
				assert.Nil(t, err)
				assert.Equal(t, args.expectedUser, res)
			},
		},
		{
			name: "Match email",
			args: args{
				expectedUser: &model.User{
					ID:       primitive.NewObjectID(),
					Username: "user",
					Email:    "user@gmail.com",
					Password: "hashed",
				},
			},
			expectations: func(mt *mtest.T, args args) (res *model.User, err error) {
				s := store.NewStore(mt.DB)
				mt.AddMockResponses(mtest.CreateCursorResponse(1, "test.test", mtest.FirstBatch, bson.D{
					{"_id", args.expectedUser.ID},
					{"username", args.expectedUser.Username},
					{"email", args.expectedUser.Email},
					{"password", args.expectedUser.Password},
				}))

				return s.FindUserByUsernameOrEmail(context.TODO(), "not username", args.expectedUser.Email)
			},
			result: func(args args, res *model.User, err error) {
				assert.Nil(t, err)
				assert.Equal(t, args.expectedUser, res)
			},
		},
		{
			name: "No match",
			args: args{
				expectedUser: &model.User{
					ID:       primitive.NewObjectID(),
					Username: "user",
					Email:    "user@gmail.com",
					Password: "hashed",
				},
			},
			expectations: func(mt *mtest.T, args args) (res *model.User, err error) {
				s := store.NewStore(mt.DB)
				mt.AddMockResponses(mtest.CreateCursorResponse(0, "test.test", mtest.FirstBatch))

				return s.FindUserByUsernameOrEmail(context.TODO(), args.expectedUser.Username, args.expectedUser.Email)
			},
			result: func(args args, res *model.User, err error) {
				assert.Equal(t, mongo.ErrNoDocuments, err)
			},
		},
	}

	for _, tc := range testCases {
		mt.Run(tc.name, func(mt *mtest.T) {
			res, err := tc.expectations(mt, tc.args)
			tc.result(tc.args, res, err)
		})
	}
}

func TestFindUserByUsernameOrEmail(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	type args struct {
		expectedUser *model.User
	}
	testCases := []struct {
		name         string
		args         args
		expectations func(mt *mtest.T, args args) (err error)
		result       func(args args, err error)
	}{
		{
			name: "Success insert",
			args: args{
				expectedUser: &model.User{
					ID:       primitive.NewObjectID(),
					Username: "user",
					Email:    "user@gmail.com",
					Password: "hashed",
				},
			},
			expectations: func(mt *mtest.T, args args) (err error) {
				s := store.NewStore(mt.DB)
				mt.AddMockResponses(mtest.CreateSuccessResponse())

				return s.InsertOne(context.TODO(), args.expectedUser)
			},
			result: func(args args, err error) {
				assert.Nil(t, err)
			},
		},
		{
			name: "Error while insert",
			args: args{
				expectedUser: &model.User{
					ID:       primitive.NewObjectID(),
					Username: "user",
					Email:    "user@gmail.com",
					Password: "hashed",
				},
			},
			expectations: func(mt *mtest.T, args args) (err error) {
				s := store.NewStore(mt.DB)
				mt.AddMockResponses(bson.D{{"not ok", 0}})

				return s.InsertOne(context.TODO(), args.expectedUser)
			},
			result: func(args args, err error) {
				assert.NotNil(t, err)
			},
		},
	}

	for _, tc := range testCases {
		mt.Run(tc.name, func(mt *mtest.T) {
			err := tc.expectations(mt, tc.args)
			tc.result(tc.args, err)
		})
	}
}
