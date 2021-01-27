package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type User struct {
	Name string
	Age  int
}

func main() {
	// Context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// DB Client
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(ctx)

	// Ping
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println("connection error:", err)
	} else {
		fmt.Println("connection success:")
	}

	// Database
	db := client.Database("mydb")
	// Collection
	usersCollection := db.Collection(("users"))

	// INSERT
	data := User{Name: "dev-sota", Age: 25}
	res, _ := usersCollection.InsertOne(ctx, data)
	fmt.Println(res)

	// READ ALL
	cursor, err := usersCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	var users []User
	_ = cursor.All(ctx, &users)
	fmt.Println(users)
}
