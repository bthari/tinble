package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	Username         string             `bson:"username"`
	Email            string             `bson:"email"`
	Password         string             `bson:"password"`
	SubscriptionType int                `bson:"subscription_type"`
	Information      *Information       `bson:"information,omitempty"`
	Preference       *Preference        `bson:"preference,omitempty"`
}

type Information struct {
	Name string `bson:"name"`
	Age  int    `bson:"age"`
}

type Preference struct {
	AgeMax int `bson:"age_max"`
	AgeMin int `bson:"age_min"`
}
