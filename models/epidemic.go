package models

import (
    "context"
    "log"
    
    "go.mongodb.org/mongo-driver/bson"

    db "nthu-chatbot-api/database"
)

type Epidemic struct {
    ID        int32     `json:"id" form:"id"`
    Category  string    `json:"category" form:"category"`
    Title     string    `json:"title" form:"title"`
    Content   []string  `json:"content" form:"content"`
}

func (e *Epidemic) GetOne(category string) (content []string, err error) {
    epidemicCollect := db.MongoDatabase.Collection("epidemic_data")

    var result Epidemic
    err = epidemicCollect.FindOne(context.TODO(), bson.M{"category": category}).Decode(&result)
    if err != nil {
        log.Println(err)
    }

    content = result.Content
    return
}

func (e *Epidemic) GetAll() (data []Epidemic, err error) {
    data = make([]Epidemic, 0)

    epidemicCollect := db.MongoDatabase.Collection("epidemic_data")
    cur, err := epidemicCollect.Find(context.TODO(), bson.M{})

    for cur.Next(context.TODO()) {
        var elem Epidemic
        err := cur.Decode(&elem)
        if err != nil {
            log.Fatal(err)
        }
        data = append(data, elem)
    }

    return
}

func (e *Epidemic) UpdateOne() (err error) {
    epidemicCollect := db.MongoDatabase.Collection("epidemic_data")

    filter := bson.M{"category": e.Category}
    option := bson.M{
        "$set": bson.M{
            "content": e.Content,
        },
    }
    _, err = epidemicCollect.UpdateOne(context.TODO(), filter, option)
    if err != nil {
        log.Println(err)
    }

    return
}
