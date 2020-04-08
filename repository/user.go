package repository

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	"bitbucket.org/suthisakch/rentala/model"
	"golang.org/x/crypto/bcrypt"
)

const (
	userConlection = "user"
)

//UserRepositoryMongo handler struct
type UserRepositoryMongo struct {
	ConnectionDB *mongo.Database
}

//UserRepository interface class
type UserRepository interface {
	GetUser(email string, pass string) (model.User, error)
}

//GetUser by Email and Password
func (userMongo *UserRepositoryMongo)GetUser(email string, pass string) (model.User, error){
	var user model.User
	//pass := utils.HashAndSalt([]byte(password))
	filter := bson.M{"email": email}
	err := userMongo.ConnectionDB.Collection(userConlection).FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		log.Println(err)
	}
	err2 := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))
	if err2 != nil {
		return user, err2
	}
	return user, err
}