package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gitlab.com/nisarg/7/controllers"
	"gitlab.com/nisarg/7/services"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server         *gin.Engine
	userservice    services.UserService
	usercontroller controllers.UserController
	ctx            context.Context
	usercollection *mongo.Collection
	mongoclient    *mongo.Client
	err            error
)

func init() {
	ctx = context.TODO()
	mongoconn := options.Client().ApplyURI("mongodb://root:smart@localhost:27017")
	mongoclient, err = mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal("err in db connection", err)
	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("err in db ping", err)
	}
	fmt.Println("mongo connection established !!!")

	usercollection = mongoclient.Database("userdb").Collection("users")
	userservice = services.NewUserService(usercollection, ctx)
	usercontroller = controllers.New(userservice)
	server = gin.Default()

}

func main() {
	defer mongoclient.Disconnect(ctx)
	basepath := server.Group("/v1")

	usercontroller.RegisterUserRoutes(basepath)
	log.Fatal(server.Run(":9000"))

}
