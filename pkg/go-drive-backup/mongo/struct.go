package mongo

import "go.mongodb.org/mongo-driver/bson/primitive"

type ObjId struct {
    Id primitive.ObjectID `json:"$oid" bson:"$oid`
}

/* user */
type User struct {
    TempObjId  primitive.ObjectID  `json:"-" bson:"_id"`
    ObjId      ObjId               `json:"_id"`

    UserID    string    `json:"userID" form:"userID"`
    Category  string    `json:"userID" form:"category"`
    Time      string    `json:"time" form:"time"`
    Flag      string    `json:"flag" form:"flag"`
    Tag       []string  `json:"tag" form:"tag"`
    StudentID string    `json:"studentID" form:"studentID"`
}

/* user map record */
type UserMapRecord struct {
    TempObjId  primitive.ObjectID  `json:"-" bson:"_id"`
    ObjId      ObjId               `json:"_id"`
    
    UserID string   `json:"userID"  bson:"userID"`
    Record []Record `json:"record"  bson:"record"`
}

type Record struct {
    Location string `json:"location" bson:"location"`
    Times    int32  `json:"times"    bson:"times"`
}

/* admin */
type Admin struct {
    TempObjId  primitive.ObjectID  `json:"-" bson:"_id"`
    ObjId      ObjId               `json:"_id"`

    Account    string              `json:"account" bson:"account"`
    Password   string              `json:"password" bson:"password"`
    Department string              `json:"department" bson:"department"`
    Name       string              `json:"name" bson:"name"`
}

/* bus schedule */
type BusSched struct {
    TempObjId  primitive.ObjectID  `json:"-" bson:"_id"`
    ObjId      ObjId               `json:"_id"`
    
    BusType    string    `json:"busType" form:"busType"`
    WdType     string    `json:"wdType" form:"wdType"`
    Direction  string    `json:"direction" form:"direction"`
    Line       string    `json:"line" form:"line"`
    DepStop    string    `json:"depStop" form:"depStop"`
    Route      []string  `json:"route" form:"route"`
    Schedule   []string  `json:"schedule" form:"schedule"`
}

/* anecdote */
type Andx struct {
    TempObjId  primitive.ObjectID  `json:"-" bson:"_id"`
    ObjId      ObjId               `json:"_id"`

    ID      int32                  `json:"id" bson:"id"`
    UserID  string                 `json:"userID" bson:"userID"`
    Time    string                 `json:"time" bson:"time"`
    Content string                 `json:"content" bson:"content"`
}

/* feedback */
type Feedback struct {
    TempObjId  primitive.ObjectID  `json:"-" bson:"_id"`
    ObjId      ObjId               `json:"_id"`

    ID        int32                `json:"id" bson:"id"`
    UserID    string               `json:"userID" bson:"userID"`
    Category  string               `json:"category" bson:"category"`
    Time      string               `json:"time" bson:"time"`
    Content   string               `json:"content" bson:"content"`
}

/* instanews */
type News struct {
    TempObjId  primitive.ObjectID  `json:"-" bson:"_id"`
    ObjId      ObjId               `json:"_id"`
    
    ID        int32    `json:"id" form:"id"`
    PostTime  string   `json:"postTime" form:"postTime"`
    Category  string   `json:"category" form:"category"`
    Title     string   `json:"title" form:"title"`
    Date      []string `json:"date" form:"date"`
    Time      string   `json:"time" form:"time"`
    Location  string   `json:"location" form:"location"`
    SourceUrl string   `json:"sourceUrl" form:"sourceUrl"`
    ImgUrl    string   `json:"imgUrl" form:"imgUrl"`
    Content   string   `json:"content" form:"content"`
}

/* qa_data */
type QA struct {
    TempObjId  primitive.ObjectID  `json:"-" bson:"_id"`
    ObjId      ObjId               `json:"_id"`

    ID        int32                `json:"id" bson:"id"`
    Category  string               `json:"category" bson:"category"`
    Data      []QAData             `json:"data" bson:"data"`
}

type QAData struct {
    Ques string `json:"ques" bson:"ques"`
    Ans  string `json:"ans" bson:"ans"`
}

/* token */
type Token struct {
    TempObjId  primitive.ObjectID  `json:"-" bson:"_id"`
    ObjId      ObjId               `json:"_id"`

    ID         string              `json:"id" bson:"id"`
    Token      string              `json:"token" bson:"token"`
}