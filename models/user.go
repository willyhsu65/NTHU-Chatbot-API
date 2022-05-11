package models

import (
    "context"
    "fmt"
    "log"
    "time"
    "errors"
    "sort"
    
    "go.mongodb.org/mongo-driver/bson"

    db "nthu-chatbot-api/database"
    "nthu-chatbot-api/utils/collectionUtil"
    "nthu-chatbot-api/utils/baseUtil"
    "nthu-chatbot-api/vendors/mongo"
)

type User struct {
    UserID          string    `json:"userID" form:"userID" bson:"userID"`
    Category        string    `json:"userID" form:"category" bson:"category"`
    Time            string    `json:"time" form:"time" bson:"time"`
    Flag            string    `json:"flag" form:"flag" bson:"flag"`
    Tag             []string  `json:"tag" form:"tag" bson:"tag"`
    StudentID       string    `json:"studentID" form:"studentID" bson:"studentID"`
    BroadcastTag    int       `json:"broadcastTag" form:"broadcastTag" bson:"broadcastTag"`
}

type UserInfo struct {
    UserID string    `json:"userID" form:"userID" bson:"userID"`
    Time   string    `json:"time" form:"time" bson:"time"`
    Tag    []string  `json:"tag" form:"tag" bson:"tag"`
}

func (u *User) InsertOne() (err error) {
    t := time.Now()
    time := t.Format("2006-01-02 15:04:05")
    category := "official"

    userCollect := db.MongoDatabase.Collection("user")

    // 找是否有重複的id
    isExist := collectionUtil.IsUserExistInCollection(userCollect, u.UserID)
    if isExist {
        log.Println("The userID is already exist:", u.UserID)
        return
    }

    // 沒有重複則新增
    u.Category = category
    u.Time = time
    u.Flag = "init"
    u.Tag = make([]string, 0)
    u.StudentID = ""

    _, err = userCollect.InsertOne(context.TODO(), u)
    if err != nil {
        log.Println(err.Error())
    }
    return
}

func (u *User) IsExist() (isExist bool, err error) {
    var result User

    userCollect := db.MongoDatabase.Collection("user")
    filter := bson.M{"userID": u.UserID}
    ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
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

func (u *User) GetFlag() (flag string, err error) {
    // FIXME: 暫時強制新增user, 捕捉沒有新增到的user
    u.InsertOne()
    
    var result User

    userCollect := db.MongoDatabase.Collection("user")
    filter := bson.M{"userID": u.UserID}
    err = userCollect.FindOne(context.TODO(), filter).Decode(&result)
    if err != nil {
        log.Println(err.Error())
    } else {
        flag = result.Flag
    }
    return
}

func (u *User) InitFlag() (err error) {
    userCollect := db.MongoDatabase.Collection("user")

    filter := bson.M{"userID": u.UserID}
    update := bson.M{
        "$set": bson.M{
            "flag": "init",
        },
    }

    ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
    _, err = userCollect.UpdateOne(ctx, filter, update)
    if err != nil {
        log.Println(err.Error())
    }
    return
}

func (u *User) SetFlag(flag string) (err error) {
    userCollect := db.MongoDatabase.Collection("user")
    
    filter := bson.M{"userID": u.UserID}
    update := bson.M{ 
        "$set": bson.M{
            "flag": flag,
        },
    }

    ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
    _, err = userCollect.UpdateOne(ctx, filter, update)
    if err != nil {
        log.Println(err.Error())
    }
    return
}

func (u *User) GetInfo() (time string, err error) {
    var result User

    userCollect := db.MongoDatabase.Collection("user")
    filter := bson.M{"userID": u.UserID}
    err = userCollect.FindOne(context.TODO(), filter).Decode(&result)
    if err != nil {
        log.Println(err.Error())
    } else {
        time = result.Time
    }
    return
}

func (u *User) UpdateBroadcastTag(tag int) (err error) {
    userCollect := db.MongoDatabase.Collection("user")
    
    filter := bson.M{"userID": u.UserID}
    update := bson.M{ "$set": bson.M{ "broadcastTag": tag } }

    ctx, _ := context.WithTimeout(context.TODO(), 5*time.Second)
    _, err = userCollect.UpdateOne(ctx, filter, update)
    if err != nil {
        log.Println(err.Error())
    }
    return
}

func (u *User) GetBroadcastAudienceIds(id string) (err error, audienceIds string) {    
    var results []User

    ts := time.Now().Unix()
    userCollect := db.MongoDatabase.Collection("user")
    filter := bson.M{"broadcastTag": bson.M{ "$gte": ts } }

    cursor, err_find := userCollect.Find(context.TODO(), filter)
    if err_find != nil {
        log.Println(err_find.Error())
        return
    }

    err := cursor.All(&results)
    if err != nil {
        log.Println(err.Error())
        return
    }

    audienceIds = ""
    for _, user_row := range results {
        if id != "" {
            if user_row.UserID == id {
                audienceIds = audienceIds + "," + user_row.UserID
            }
        } else {
            audienceIds = audienceIds + "," + user_row.UserID
        }
    }
    return
}

/* user_map_record */
type UserMapRecord struct {
    UserID string   `json:"userID"  bson:"userID"`
    Record []Record `json:"record"  bson:"record"`
}

type Record struct {
    Location string `json:"location" bson:"location"`
    Times    int32  `json:"times"    bson:"times"`
}

func (u *User) InsertMapRecord(location string) (err error) {
    userMapRecordCollect := db.MongoDatabase.Collection("user_map_record")

    /* 檢查是否有此userID的紀錄, 沒有則新增 */
    var result UserMapRecord
    var addFlag bool = false
    err = userMapRecordCollect.FindOne(context.TODO(), bson.M{"userID": u.UserID}).Decode(&result)
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
        c := UserMapRecord{
            UserID: u.UserID,
            Record: make([]Record, 0),
        }
    
        _, err = userMapRecordCollect.InsertOne(context.TODO(), c)
        if err != nil{
            log.Println(err)
        } else {
            log.Println("Successful insert one user map record")
        }
    } else {
        err = errors.New(fmt.Sprintf("The '%s' userID is exist, reject this add action.", u.UserID))
    }

    /* 紀錄location */
    // 已有此 location times+1
    var locationFlag bool = false
    var tempRecord []Record
    for _, record := range result.Record {
        if record.Location == location {
            record.Times += 1
            locationFlag = true
        }
        tempRecord = append(tempRecord, record)
    }
    result.Record = tempRecord

    // 沒有此 location, 則新增
    if !locationFlag {
        newLocationRecord := Record{
            Location: location,
            Times: 1,
        }

        result.Record = append(result.Record, newLocationRecord)
    }

    // 更新 user's map record
    filter := bson.M{"userID": u.UserID}
    option := bson.M{
        "$set": bson.M{
            "record": result.Record,
        },
    }
    _, err = userMapRecordCollect.UpdateOne(context.TODO(), filter, option)
    if err != nil {
        log.Println(err)
    }
    return
}

func (u *User) GetMapRecord() (locations []string, err error) {
    locations = make([]string, 0)

    userMapRecordCollect := db.MongoDatabase.Collection("user_map_record")

    var addFlag bool = false
    var result UserMapRecord
    err = userMapRecordCollect.FindOne(context.TODO(), bson.M{"userID": u.UserID}).Decode(&result)
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
        c := UserMapRecord{
            UserID: u.UserID,
            Record: make([]Record, 0),
        }
    
        _, err = userMapRecordCollect.InsertOne(context.TODO(), c)
        if err != nil{
            log.Println(err)
        } else {
            log.Println("Successful insert one user map record")
            return
        }
    }

    // 依照 times 由大至小排序地點
    record := result.Record
    sort.Slice(record,  func(i, j int) bool { 
        return record[i].Times > record[j].Times 
    })

    for _, r := range record {
        locations = append(locations, r.Location)
    }

    // 如果超過10則, 只回前面10個
    if len(locations) > 10 {
        locations = locations[:10]
    }

    return
}


// audience tag add in user's tag category
func (u *User) UpdateTag(userID string, audienceDescription string) (err error) {
    userCollect := db.MongoDatabase.Collection("user")

    filter := bson.M{"userID": userID}
    option := bson.M{
        "$addToSet": bson.M{
            "tag": bson.M{
                "$each": []string{audienceDescription},
            },
        },
    }
    _, err = userCollect.UpdateOne(context.TODO(), filter, option)
    if err != nil {
        log.Println(err)
    }
    return
}

/*
* 刪除audience後, 更新所有user的tag欄位
* @param audienceDescription string
* @return count int 成功刪除幾個人的tag
* @return err error
*/
func (u *User) RemoveTagByAudiDesc(audienceDescription string) (count int, err error) {
    // get all user info including tag category
    userCollect := db.MongoDatabase.Collection("user")

    userInfoSlice := make([]UserInfo, 0)
    cur, err := userCollect.Find(context.TODO(), bson.M{"category": "official"})
    if err != nil { log.Println(err) }

    for cur.Next(context.TODO()) {
        var result UserInfo
        err := cur.Decode(&result)
        if err != nil { log.Println(err) }

        userInfoSlice = append(userInfoSlice, result)
    }
    if err := cur.Err(); err != nil {
        log.Println(err)
    }

    // if include audienceDescription, remove tag from user's tag category
    for _, userInfo := range(userInfoSlice) {
        var isContainTag bool = false
        if baseUtil.Contains(userInfo.Tag, audienceDescription) {
            isContainTag = true
        }

        // 刪除tag後，更新tag欄位
        if isContainTag {
            var result UserInfo
            filter := bson.M{"userID": userInfo.UserID}
            err = userCollect.FindOne(context.TODO(), filter).Decode(&result)
            if err != nil {
                log.Println(err.Error())
                return
            }
            count = count + 1

            // 更新 tag
            newTagSlice := baseUtil.DeleteEle(result.Tag, audienceDescription)
            filter = bson.M{"userID": userInfo.UserID}
            option := bson.M{
                "$set": bson.M{
                    "tag": newTagSlice,
                },
            }
            _, err = userCollect.UpdateOne(context.TODO(), filter, option)
            if err != nil {
                log.Println(err)
                return
            }
        }
    }
    return
}


/*
* 確認是否已經綁定學號
* @param userID string
* @param studentID string
* @return err error
*/
func (u *User) CheckBindedStudentID() (isBind bool, studentID string, err error) {
    userCollect := db.MongoDatabase.Collection("user")

    var result User
    err = userCollect.FindOne(context.TODO(), bson.M{"userID": u.UserID}).Decode(&result)
    if err != nil {
        log.Println(err.Error())
        return
    }

    if result.StudentID != "" {
        isBind = true
        studentID = result.StudentID
    }

    return
}


/*
* 綁定學生學號
* @param userID string
* @param studentID string
* @return err error
*/
func (u *User) BindStudentID(studentID string) (err error) {
    userCollect := db.MongoDatabase.Collection("user")

    filter := bson.M{"userID": u.UserID}
    option := bson.M{
        "$set": bson.M{
            "studentID": studentID,
        },
    }
    _, err = userCollect.UpdateOne(context.TODO(), filter, option)

    return
}