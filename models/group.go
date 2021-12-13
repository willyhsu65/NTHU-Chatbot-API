package models

import (
    "context"
    "log"
    "time"
    
    "go.mongodb.org/mongo-driver/bson"

    db "nthu-chatbot-api/database"
)

type Group struct {
    Account string
}

/* anecdote */
func (g *Group) GetAndx() (andxSlice []Andx, err error) {
    andxSlice = make([]Andx, 0)

    anecdoteCollect := db.MongoDatabase.Collection("anecdote")
    ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

    cur, err := anecdoteCollect.Find(ctx, bson.D{})
    if err != nil {
        log.Println(err) 
    }

    for cur.Next(ctx) {
        var result Andx
        err := cur.Decode(&result)
        if err != nil { log.Println(err) }

        andxSlice = append(andxSlice, result)
    }
    if err := cur.Err(); err != nil {
        log.Println(err)
    }

    return
}

func (g *Group) DeleteAndx(id int32) (err error) {
    anecdoteCollect := db.MongoDatabase.Collection("anecdote")
    ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

    _, err = anecdoteCollect.DeleteOne(ctx, bson.M{"id": id})
    if err != nil {
        log.Println(err)
    }
    return
}

/* feedback */
func (g *Group) GetFeedback(category string) (fbSlice []Feedback, err error) {
    fbSlice = make([]Feedback, 0)

    feedbackCollect := db.MongoDatabase.Collection("feedback")
    ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

    cur, err := feedbackCollect.Find(ctx, bson.M{"category": category})
    if err != nil {
        log.Println(err) 
    }

    for cur.Next(ctx) {
        var result Feedback
        err := cur.Decode(&result)
        if err != nil { log.Println(err) }

        fbSlice = append(fbSlice, result)
    }
    if err := cur.Err(); err != nil {
        log.Println(err)
    }

    return
}

func (g *Group) DeleteFeedback(id int32) (err error) {
    feedbackCollect := db.MongoDatabase.Collection("feedback")
    ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

    _, err = feedbackCollect.DeleteOne(ctx, bson.M{"id": id})
    if err != nil {
        log.Println(err)
    }
    return
}