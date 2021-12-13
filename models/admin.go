package models

import (
    "context"
    _"fmt"
    "strings"
    "log"
    "time"
    
    "go.mongodb.org/mongo-driver/bson"

    db "nthu-chatbot-api/database"
    "nthu-chatbot-api/vendors/mongo"
)

type Admin struct {
    Account    string
    Password   string
    Department string
    Name       string
}

func (a *Admin) UserDetail() (result Admin, isExist bool, err error) {
    userCollect := db.MongoDatabase.Collection("admin")
    ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

    filter := bson.M{"account": a.Account, "password": a.Password}
    err = userCollect.FindOne(ctx, filter).Decode(&result)
    if err != nil {
        log.Println(err)
        if err.Error() == mongo.ErrNoDocument {
            isExist = false
            err = nil
        }
    } else {
        isExist = true
    }

    return
}

func (a *Admin) GetUserInfo() (userInfoSlice []UserInfo, err error) {
    userInfoSlice = make([]UserInfo, 0)

    userCollect := db.MongoDatabase.Collection("user")
    ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

    cur, err := userCollect.Find(ctx, bson.M{"category": "official"})
    if err != nil { log.Println(err) }

    for cur.Next(ctx) {
        var result UserInfo
        err := cur.Decode(&result)
        if err != nil { log.Println(err) }

        userInfoSlice = append(userInfoSlice, result)
    }
    if err := cur.Err(); err != nil {
        log.Println(err)
    }

    return
}

func (a *Admin) GetUserNewCount() (todayCount int, err error) {
    todayCount = 0

    userCollect := db.MongoDatabase.Collection("user")
    ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

    cur, err := userCollect.Find(ctx, bson.M{"category": "official"})
    if err != nil { log.Println(err) }

    var userInfoSlice []UserInfo
    for cur.Next(ctx) {
        var result UserInfo
        err := cur.Decode(&result)
        if err != nil { log.Println(err) }

        userInfoSlice = append(userInfoSlice, result)
    }
    if err := cur.Err(); err != nil {
        log.Println(err)
    }

    
    t := time.Now()
    todayDate := t.Format("2006-01-02")

    for _, user := range userInfoSlice {
        signsUpDate := strings.Split(user.Time, " ")[0]
        
        if signsUpDate == todayDate {
            todayCount++;
        }
    }
    return
}

func (a *Admin) GetBetaUserInfo() (userInfoSlice []UserInfo, err error) {
    userInfoSlice = make([]UserInfo, 0)
    
    userCollect := db.MongoDatabase.Collection("user")
    ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

    cur, err := userCollect.Find(ctx, bson.M{"category": "beta"})
    if err != nil { log.Println(err) }

    for cur.Next(ctx) {
        var result UserInfo
        err := cur.Decode(&result)
        if err != nil { log.Println(err) }

        userInfoSlice = append(userInfoSlice, result)
    }
    if err := cur.Err(); err != nil {
        log.Println(err)
    }

    return
}