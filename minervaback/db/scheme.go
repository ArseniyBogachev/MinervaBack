package db

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	// System fields
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Login    string             `json:"login" bson:"login"`
	Password string             `json:"-" bson:"password"`
	IsAdmin  bool               `json:"is_admin" bson:"is_admin"`

	// User fields
	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" bson:"last_name"`
	Email     string `json:"email" bson:"email"`
	Phone     string `json:"phone" bson:"phone"`
}

type Form struct {
	// System fields
	ID      primitive.ObjectID `json:"id" bson:"_id"`
	OwnerID primitive.ObjectID `json:"owner_id" bson:"owner_id"`

	// Form fields
	Title string `json:"title" bson:"title"`
}

type FormBlock struct {
	// System fields
	ID     primitive.ObjectID `json:"id" bson:"_id"`
	FormID primitive.ObjectID `json:"form_id" bson:"-"`

	// Form block fields
	Order int         `json:"order" bson:"order"`
	Data  interface{} `json:"data" bson:"data"`
}
