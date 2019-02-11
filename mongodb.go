package main

import (
	"context"
	"fmt"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

var Client mongo.Client
var Database *mongo.Database
var Ctx context.Context

func ConnectMongo() {
	Client, err := mongo.NewClient("mongodb://localhost:27017")
	if err != nil {
		fmt.Println("Not able to connect to mongodb")
		return
	}
	Ctx, _ := context.WithTimeout(context.Background(), 10*time.Minute)
	err = Client.Connect(Ctx)
	if err != nil {
		fmt.Println("Connection error", err.Error())
	}
	Database := Client.Database("todo")
	if Database != nil {

	}
	Database.Collection("user").InsertOne(Ctx, bson.M{"user": "mathan"})
}

func SaveUser(user *User) error {
	//ctx, _ := context.WithTimeout(context.Background(), 10*time.Minute)
	if Database.Client == nil {
		fmt.Println("Client nil")
		return nil
	}

	if user == nil {
		fmt.Println("User nil")
	}
	Database := Client.Database("todo")
	collection := Database.Collection("user")
	res, err := collection.InsertOne(Ctx, bson.M{"user": "guru"})
	if err != nil {
		return err
	}
	fmt.Println("Id: ", res.InsertedID)
	return nil
}
