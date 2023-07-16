package main

import (
	"context"
	"fmt"
	"golangGin/controllers"
	"golangGin/service"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server         *gin.Engine
	userService    service.UserService
	userController controllers.UserController
	ctx            context.Context
	userCollection *mongo.Collection
	mongoClient    *mongo.Client
	err            error
)

func init() {
	ctx = context.TODO()
	mongoConn := options.Client().ApplyURI("mongodb://localhost:27017")
	mongoClient, err := mongo.Connect(ctx, mongoConn)
	if err != nil {
		log.Fatal(err)
	}
	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB connection is successfull!")

	userCollection = mongoClient.Database("userdb").Collection("users")
	userService = service.NewUserService(userCollection, ctx)
	userController = controllers.NewUserController(userService)
	server = gin.Default()
}

func main() {
	defer mongoClient.Disconnect(ctx)

	basePath := server.Group("/v1")
	userController.RegisterUserRoutes(basePath)

	log.Fatal(server.Run(":8080"))
}
