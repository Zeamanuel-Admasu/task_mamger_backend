package db

import (
	"context"
	"errors"

	"github.com/zaahidali/task_manager_api/internal/domain/models"
	"github.com/zaahidali/task_manager_api/internal/domain/ports"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type MongoUserRepo struct {
	Collection *mongo.Collection
}

func NewMongoUserRepo(col *mongo.Collection) ports.UserRepository {
	return &MongoUserRepo{Collection: col}
}

func (r *MongoUserRepo) CreateUser(ctx context.Context, user models.User) error {
	user.ID = primitive.NewObjectID()
	_, err := r.Collection.InsertOne(ctx, user)
	return err
}

func (r *MongoUserRepo) AuthenticateUser(ctx context.Context, username, password string) (*models.User, error) {
	var user models.User
	err := r.Collection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid credentials")
	}
	return &user, nil
}

func (r *MongoUserRepo) PromoteUser(ctx context.Context, username string) error {
	result, err := r.Collection.UpdateOne(ctx, bson.M{"username": username}, bson.M{"$set": bson.M{"role": "admin"}})
	if err != nil || result.MatchedCount == 0 {
		return errors.New("user not found or promotion failed")
	}
	return nil
}

func (r *MongoUserRepo) Count(ctx context.Context) (int64, error) {
	return r.Collection.CountDocuments(ctx, bson.D{})
}
