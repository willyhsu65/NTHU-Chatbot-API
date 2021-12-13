package models

import (
    "context"
    "log"
    
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson"

    db "nthu-chatbot-api/database"
)

type Phone struct {
    Name  string `json:"name" bson:"name"`
    Phone string `json:"phone" bson:"phone"`
}

func (p *Phone) GetPhones() (phoneSlice []Phone, err error) {
    phoneSlice = make([]Phone, 0)

    AcadPhoneCollect := db.MongoDatabase.Collection("academic_phone")
    acadCur, err := AcadPhoneCollect.Find(context.TODO(), bson.M{})
    if err != nil {
        log.Println("Get academic phones error:", err)
        return 
    }
    for acadCur.Next(context.TODO()) {
        var result Phone
        err := acadCur.Decode(&result)
        if err != nil { 
            log.Println(err) 
        }

        phoneSlice = append(phoneSlice, result)
    }
    if err := acadCur.Err(); err != nil {
        log.Println(err)
    }

    adminPhoneCollect := db.MongoDatabase.Collection("administration_phone")
    adminCur, err := adminPhoneCollect.Find(context.TODO(), bson.M{})
    if err != nil {
        log.Println("Get administration phones error:", err)
        return 
    }
    for adminCur.Next(context.TODO()) {
        var result Phone
        err := adminCur.Decode(&result)
        if err != nil { 
            log.Println(err) 
        }

        phoneSlice = append(phoneSlice, result)
    }
    if err := adminCur.Err(); err != nil {
        log.Println(err)
    }

    return
}

func (p *Phone) GetPhonesByType(phone_type string) (phoneSlice []Phone, err error) {
    phoneSlice = make([]Phone, 0)

    var phoneCollect *mongo.Collection
    if phone_type == "academic" {
        phoneCollect = db.MongoDatabase.Collection("academic_phone")
    } else if phone_type == "administration" {
        phoneCollect = db.MongoDatabase.Collection("administration_phone")
    }

    cur, err := phoneCollect.Find(context.TODO(), bson.M{})
    if err != nil {
        log.Println("Get phones error:", err)
        return 
    }

    for cur.Next(context.TODO()) {
        var result Phone
        err := cur.Decode(&result)
        if err != nil { 
            log.Println(err) 
        }

        phoneSlice = append(phoneSlice, result)
    }
    if err := cur.Err(); err != nil {
        log.Println(err)
    }

    return
}
