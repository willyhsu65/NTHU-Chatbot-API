package models

import (
    "context"
    "fmt"
    "log"
    "errors"
    
    "go.mongodb.org/mongo-driver/bson"

    db "nthu-chatbot-api/database"
    "nthu-chatbot-api/utils/collectionUtil"
    "nthu-chatbot-api/vendors/mongo"
)

type QAInfo struct {
    Category  string `json:"category" form:"category"`
    Ques string `json:"ques" form:"ques"`
    Ans  string `json:"ans" form:"ans"`
}

type QAData struct {
    Ques string `json:"ques" form:"ques"`
    Ans  string `json:"ans" form:"ans"`
}

type QA struct {
    ID        int32     `json:"id" form:"id"`
    Category  string    `json:"category" form:"category"`
    Data      []QAData  `json:"data" form:"data"`
}

func (qa *QA) AddCategory(category string) (err error) {
    qaCollect := db.MongoDatabase.Collection("qa_data")

    id := collectionUtil.GetCollectionLastId(qaCollect)

    // 看是否已經存在此category
    var result QA
    var addFlag bool = false
    err = qaCollect.FindOne(context.TODO(), bson.M{"category": category}).Decode(&result)
    if err != nil {
        if err.Error() == mongo.ErrNoDocument {
            addFlag = true
        } else {
            log.Println(err)
            return
        }
    }

    // 新增
    if addFlag {
        c := QA{
            ID: id,
            Category: category,
            Data: make([]QAData, 0),
        }
    
        _, err = qaCollect.InsertOne(context.TODO(), c)
        if err != nil{
            log.Println(err)
        } else {
            log.Println("Successful add qa category")
        }
    } else {
        err = errors.New(fmt.Sprintf("The '%s' category is exist, reject this add action.", category))
    }

    return
}

func (qa *QA) AddOne(qaData QAData) (err error) {
    qaCollect := db.MongoDatabase.Collection("qa_data")
    
    filter := bson.M{"category": qa.Category}
    option := bson.M{
        "$push": bson.M{
            "data": qaData,
        },
    }
    _, err = qaCollect.UpdateOne(context.TODO(), filter, option)
    if err != nil {
        log.Println(err)
    }
    return
}

func (qa *QA) GetOneData(category string) (qaData []QAData, err error) {
    qaData = make([]QAData, 0)

    qaCollect := db.MongoDatabase.Collection("qa_data")

    var result QA
    err = qaCollect.FindOne(context.TODO(), bson.M{"category": category}).Decode(&result)
    if err != nil {
        log.Println(err)
    }

    qaData = result.Data
    return
}

func (qa *QA) GetAllData() (qaSlice []QA, err error) {
    qaSlice = make([]QA, 0)

    qaCollect := db.MongoDatabase.Collection("qa_data")
    cur, err := qaCollect.Find(context.TODO(), bson.M{})

    for cur.Next(context.TODO()) {
        var elem QA
        err := cur.Decode(&elem)
        if err != nil {
            log.Fatal(err)
        }
        qaSlice = append(qaSlice, elem)
    }

    return
}

func (qa *QA) DelOneQA(category string, ques string) (err error) {
    qaCollect := db.MongoDatabase.Collection("qa_data")

    var result QA
    err = qaCollect.FindOne(context.TODO(), bson.M{"category": category}).Decode(&result)
    if err != nil {
        log.Println(err)
    }

    var newQaData []QAData
    qaData := result.Data
    for _, item := range qaData {
        if item.Ques != ques {
            newQaData = append(newQaData, item)
        }
    }

    // update new qaData
    filter := bson.M{"category": category}
    option := bson.M{
        "$set": bson.M{
            "data": newQaData,
        },
    }
    _, err = qaCollect.UpdateOne(context.TODO(), filter, option)
    if err != nil {
        log.Println(err)
    }

    return
}