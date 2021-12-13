package models

import (
    "context"
    _"fmt"
    "log"

    "go.mongodb.org/mongo-driver/bson"

    db "nthu-chatbot-api/database"
)

type Token struct {
    ID      string  `json:"id" form:"id"`
    Token   string  `json:"token" form:"token"`
    Webhook string  `json:"webhook" form:"webhook"`
}

func (t *Token) GetToken() (token string, err error) {
    var result Token

    tokenCollect := db.MongoDatabase.Collection("token")

    filter := bson.M{"id": t.ID}
    err = tokenCollect.FindOne(context.TODO(), filter).Decode(&result)
    if err != nil {
        log.Println("Get Token error:", err)
        return 
    }

    token = result.Token

    return
}


func (t *Token) GetTokenAndWebhook() (data Token, err error){
    var result Token

    tokenCollect := db.MongoDatabase.Collection("token")

    filter := bson.M{"id": t.ID}
    err = tokenCollect.FindOne(context.TODO(), filter).Decode(&result)
    if err != nil {
        log.Println("Get Token error:", err)
        return 
    }

    data = result
    return
}

