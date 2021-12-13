package database

import (
    "context"
    "fmt"
    "log"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
    host = "mongodb://[MONGO_ACCOUNT]:[MONGO_PASS]@[MONGO_IP]:27017"
)

var mongoClient *mongo.Client
var MongoDatabase *mongo.Database

func init() {
    var err error

    mongoURI := fmt.Sprintf(host)

    // connect to mongo
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    mongoClient, err = mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
    if err != nil {
        log.Fatal("connet err:", err.Error())
        panic(err)
    }

    // Check the connection
    ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
    err = mongoClient.Ping(ctx, readpref.Primary())
    if err != nil {
        fmt.Println("could not ping to mongo db service: %v\n", err)
        return
    } else {
        fmt.Println("Successful connected to MongoDB!")
    }

    MongoDatabase = mongoClient.Database("nthu_chatbot_db")
}