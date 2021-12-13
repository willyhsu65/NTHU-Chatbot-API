package collectionUtil

import (
	_"fmt"
	"log"
    "time"
    "reflect"
    "context"
    
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo"
)

/* @desc get collection last id
   @param collect *mongo.Collection
   @return id int32
*/
func GetCollectionLastId(collect *mongo.Collection) (id int32) {
    ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	itemCount, err := collect.CountDocuments(ctx, bson.D{})
    if err != nil {
        log.Println(err)
    }

    if int(itemCount) == 0 {
        id = 1
    } else {
        findOptions := options.Find()
        findOptions.SetLimit(1)
        findOptions.SetSort(map[string]int{"id": -1})
        
        cur, err := collect.Find(ctx, bson.M{}, findOptions)
        if err != nil {
            log.Println(err)
        }

        var s []map[string]interface{}
        for cur.Next(ctx) {
            var result map[string]interface{}
            err := cur.Decode(&result)
            if err != nil { log.Println(err) }
    
            s = append(s, result)
        }

        var lastID int32
        val := reflect.ValueOf(s[0])
        if val.Kind() == reflect.Map {
            for _, e := range val.MapKeys() {
                v := val.MapIndex(e)
                switch t := v.Interface().(type) {
                    case int32:
                        lastID = t
                }
            }
        }
        id = lastID + 1
    }

    return 
}


/* @desc 比對是否有相同的 userID 已經存在在 collection 內
   @param collect *mongo.Collection
   @param userID string
   @return isExist bool
*/
type User struct {
    UserID string `json:"userID" form:"userID"`
    Name   string `json:"name" form:"name"`
    Time   string `json:"time" form:"time"`
    Flag   string `json:"flag" form:"flag"`
}

func IsUserExistInCollection(collect *mongo.Collection, userID string) (isExist bool) {
    isExist = false

    ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

    cur, err := collect.Find(ctx, bson.M{})
    if err != nil {
        log.Println(err)
    }

    var userSlice []User
    for cur.Next(ctx) {
        var result User
        err := cur.Decode(&result)
        if err != nil { log.Println(err) }

        userSlice = append(userSlice, result)
    }

    for _, user := range userSlice {
        if user.UserID == userID {
            isExist = true
        }
    }

    return
}