package store

import (
	"context"
	"github/bthari/tinble/internal/model"
	"go.mongodb.org/mongo-driver/bson"
)

type FindByUsername struct {
	Username string `bson:"username"`
}

const (
	userCollectionName = "user"
)

func (s *Store) InsertOne(ctx context.Context, user *model.User) (err error) {
	coll := s.DB.Collection(userCollectionName)

	_, err = coll.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	return
}

func (s *Store) FindUserByUsernameOrEmail(ctx context.Context, username, email string) (user *model.User, err error) {
	coll := s.DB.Collection(userCollectionName)

	res := coll.FindOne(ctx, bson.D{
		{"$or", bson.A{
			bson.M{"username": username},
			bson.M{"email": email}},
		},
	})
	if err := res.Err(); err != nil {
		return nil, err
	}

	user = &model.User{}
	err = res.Decode(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
