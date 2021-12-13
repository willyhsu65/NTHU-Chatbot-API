package models

import (
    "context"
    "log"
    "time"

    db "nthu-chatbot-api/database"
    "nthu-chatbot-api/utils/collectionUtil"
)

type Feedback struct {
    ID        int32  `json:"id" form:"id"`
    UserID    string `json:"userID" form:"userID"`
    Category  string `json:"category" form:"category"`
    Time      string `json:"time" form:"time"`
    Content   string `json:"content" form:"content"`
}

func (f *Feedback) InsertOne(content string) (err error){
    andxCollect := db.MongoDatabase.Collection("feedback")

    id := collectionUtil.GetCollectionLastId(andxCollect)
    insertTime := time.Now().Format("2006-01-02 15:04:05")

    var fb Feedback = Feedback{
        ID: id,
        UserID: f.UserID,
        Category: f.Category,
        Time: insertTime,
        Content: content,
    }

    _, err = andxCollect.InsertOne(context.TODO(), fb)
    if err != nil{
        log.Println(err)
    } else {
        log.Println("Successful insert feedback")
    }
    return
}