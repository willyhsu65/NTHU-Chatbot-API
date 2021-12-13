package models

import (
    "context"
    _"fmt"
    _"strings"
    "log"
    "time"
    "math/rand"
    
    "go.mongodb.org/mongo-driver/bson"

    db "nthu-chatbot-api/database"
    "nthu-chatbot-api/utils/collectionUtil"
)

type Andx struct {
    ID      int32  `json:"id" form:"id"`
    UserID  string `json:"userID" form:"userID"`
    Time    string `json:"time" form:"time"`
    Content string `json:"content" form:"content"`
}

func (a *Andx) InsertOne(content string) (err error) {
    andxCollect := db.MongoDatabase.Collection("anecdote")
    
    id := collectionUtil.GetCollectionLastId(andxCollect)
    insertTime := time.Now().Format("2006-01-02 15:04:05")
    
    // insert anecdote
    var andx Andx = Andx{
        ID: id,
        UserID: a.UserID,
        Time: insertTime,
        Content: content,
    }

    _, err = andxCollect.InsertOne(context.TODO(), andx)
    if err != nil{
        log.Println(err)
    } else {
        log.Println("Successful insert anecdote")
    }
    return
}

func (a *Andx) GetOne() (content string, err error) {
    rand.Seed(time.Now().UnixNano())

    ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
    andxCollect := db.MongoDatabase.Collection("anecdote")
    cur, err := andxCollect.Find(ctx, bson.D{})

    var results []Andx
    for cur.Next(context.TODO()) {
        var elem Andx
        err := cur.Decode(&elem)
        if err != nil {
            log.Fatal(err)
        }
        results = append(results, elem)
    }

    if err := cur.Err(); err != nil {
        log.Fatal(err)
    }

    //Close the cursor once finished
    cur.Close(context.TODO())

    andxCount := len(results)
    randNum := rand.Intn(andxCount)

    content = results[randNum].Content

    return
}


