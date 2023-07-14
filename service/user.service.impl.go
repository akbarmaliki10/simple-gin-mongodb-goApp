package service

import (
	"context"
	"golangGin/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceImpl struct {
	userCollection *mongo.Collection
	ctx            context.Context
}

// make a function that act like a constructor
func NewUserService(userCollection *mongo.Collection, ctx context.Context) UserService {
	return &UserServiceImpl{
		userCollection: userCollection,
		ctx:            ctx,
	}
}

// receiver function or more like classes/struct method in python/java
func (u *UserServiceImpl) CreateUser(user *models.User) error {
	return nil
}

func (u *UserServiceImpl) GetUser(name *string) (*models.User, error) {
	return nil, nil
}

func (u *UserServiceImpl) GetAll() ([]*models.User, error) {
	return nil, nil
}

func (u *UserServiceImpl) UpdateUser(models *models.User) error {
	return nil
}

func (u *UserServiceImpl) DeleteUser(name *string) error {
	return nil
}
