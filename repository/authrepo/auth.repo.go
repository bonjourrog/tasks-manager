package authrepo

import (
	"context"
	"fmt"
	"os"

	"github.com/bonjourrog/taskm/db"
	"github.com/bonjourrog/taskm/entity"
	"github.com/bonjourrog/taskm/utils"
	"go.mongodb.org/mongo-driver/bson"
)

type AuthRepo interface {
	Register(user entity.User) entity.MongoResult
	SignIn(user entity.User) entity.MongoResult
}

type authRepo struct{}

func NewAuthRepo() AuthRepo {
	return &authRepo{}
}
func (*authRepo) Register(user entity.User) entity.MongoResult {
	var (
		_db          = db.NewMongoConnection()
		mResult      entity.MongoResult
		singleResult entity.User
	)
	client := _db.Connection()
	defer func() {
		client.Disconnect(context.TODO())
	}()
	coll := client.Database(os.Getenv("MONGO_DB")).Collection("users")
	if err := coll.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&singleResult); err != nil {
		mResult = entity.MongoResult{
			Success: false,
			Message: err.Error(),
		}
		return mResult
	}

	if singleResult.Email == user.Email {
		mResult = entity.MongoResult{
			Success: false,
			Message: "email is already registered",
		}
		return mResult
	}

	result, err := coll.InsertOne(context.TODO(), user)
	if err != nil {
		mResult = entity.MongoResult{
			Success: false,
			Message: err.Error(),
		}
		return mResult
	}
	mResult = entity.MongoResult{
		Success:    true,
		Message:    "user registered successfully",
		InsertedID: fmt.Sprintf("%v", result.InsertedID),
	}
	return mResult
}
func (*authRepo) SignIn(user entity.User) entity.MongoResult {
	var (
		_db          = db.NewMongoConnection()
		singleResult entity.User
		mResult      entity.MongoResult
	)
	client := _db.Connection()
	defer func() {
		client.Disconnect(context.TODO())
	}()
	coll := client.Database(os.Getenv("MONGO_DB")).Collection("users")
	if err := coll.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&singleResult); err != nil {
		mResult = entity.MongoResult{
			Success: false,
			Message: "invalid email or password",
		}
		return mResult
	}
	if err := utils.CheckPassword(singleResult.Password, user.Password); err != nil {
		mResult = entity.MongoResult{
			Success: false,
			Message: "invalid email or password",
		}
		return mResult
	}
	mResult = entity.MongoResult{
		Success: true,
		Message: "authentication successful",
		Data:    singleResult,
	}
	return mResult
}
