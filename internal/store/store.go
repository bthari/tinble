package store

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"

	"github/bthari/tinble/internal/model"
	"github/bthari/tinble/pkg/config"
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

func InitMongo(config *config.Database) *mongo.Database {
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
