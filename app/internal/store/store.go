package store

import (
	"context"
	"fmt"
	"github/bthari/tinble/app/internal/model"
	"github/bthari/tinble/app/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type StoreInterface interface {
	InsertOne(ctx context.Context, user *model.User) (err error)
	FindUserByUsernameOrEmail(ctx context.Context, username, email string) (user *model.User, err error)
}

type Store struct {
	DB *mongo.Database
}

func NewStore(database *mongo.Database) Store {
	return Store{
		DB: database,
	}
}

func InitStore(config *config.Database) *mongo.Database {
	credential := options.Credential{
		Username: config.User,
		Password: config.Password,
	}

	addr := fmt.Sprintf("%s:%d", config.Host, config.Port)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(addr).SetAuth(credential))
	if err != nil {
		log.Fatal(err)
	}

	return client.Database(config.Database)
}
