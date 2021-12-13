/* 備份 mongodb collection
*  匯出位置: backup/<date>
*/

package mongo

import (
    "os"
    "path"
    "log"
    "time"
    "context"
    "encoding/json"

    "go.mongodb.org/mongo-driver/bson"

    "nthu-chatbot-api/database"
    "nthu-chatbot-api/pkg/go-drive-backup"
)

func LocalBackupFunc() {
    log.Println("Start local backup mission at", time.Now())

    UserCollectionLocalBackup(goDriveBackup.DateBackupDirPath)
    UserMapRecordCollectionLocalBackup(goDriveBackup.DateBackupDirPath)
    AdminCollectionLocalBackup(goDriveBackup.DateBackupDirPath)
    BusSchedCollectionLocalBackup(goDriveBackup.DateBackupDirPath)
    AndxCollectionLocalBackup(goDriveBackup.DateBackupDirPath)
    FeedbackCollectionLocalBackup(goDriveBackup.DateBackupDirPath)
    InstanewsCollectionLocalBackup(goDriveBackup.DateBackupDirPath)
    QaDataCollectionLocalBackup(goDriveBackup.DateBackupDirPath)
    TokenCollectionLocalBackup(goDriveBackup.DateBackupDirPath)

    log.Println("Local backup mission complete at", time.Now())
}

/* mongo backup func 
* UserCollectionLocalBackup
* UserMapRecordCollectionLocalBackup
* AdminCollectionLocalBackup
* BusSchedCollectionLocalBackup
* AndxCollectionLocalBackup
* FeedbackCollectionLocalBackup
* InstanewsCollectionLocalBackup
* QaDataCollectionLocalBackup
* TokenCollectionLocalBackup
*/

func UserCollectionLocalBackup(backupDirPath string) {
    collection := database.MongoDatabase.Collection("user")

    cur, _ := collection.Find(context.TODO(), bson.M{})
    var results []User
    for cur.Next(context.TODO()) {
        var elem User
        if err := cur.Decode(&elem); err != nil {
            log.Fatal(err)
        }
        
        // set ObjId
        objId := ObjId{Id: elem.TempObjId}
        elem.ObjId = objId

        results = append(results, elem)
    }

    for _, doc := range results {
        b, err := json.Marshal(doc)
        if err != nil {
            log.Fatal(err)
        }
        
        f, err := os.OpenFile(path.Join(backupDirPath, "user.json"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
        if err != nil {
            log.Fatal(err)
        }
        if _, err := f.Write(b); err != nil {
            log.Fatal(err)
        }
        if _, err := f.Write([]byte("\n")); err != nil {
            log.Fatal(err)
        }
        if err := f.Close(); err != nil {
            log.Fatal(err)
        }
    }

    log.Println("Successfully export user collcetion at", time.Now())
}

func UserMapRecordCollectionLocalBackup(backupDirPath string) {
    collection := database.MongoDatabase.Collection("user_map_record")

    cur, _ := collection.Find(context.TODO(), bson.M{})
    var results []UserMapRecord
    for cur.Next(context.TODO()) {
        var elem UserMapRecord
        if err := cur.Decode(&elem); err != nil {
            log.Fatal(err)
        }
        
        // set ObjId
        objId := ObjId{Id: elem.TempObjId}
        elem.ObjId = objId

        results = append(results, elem)
    }

    for _, doc := range results {
        b, err := json.Marshal(doc)
        if err != nil {
            log.Fatal(err)
        }
        
        f, err := os.OpenFile(path.Join(backupDirPath, "user_map_record.json"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
        if err != nil {
            log.Fatal(err)
        }
        if _, err := f.Write(b); err != nil {
            log.Fatal(err)
        }
        if _, err := f.Write([]byte("\n")); err != nil {
            log.Fatal(err)
        }
        if err := f.Close(); err != nil {
            log.Fatal(err)
        }
    }

    log.Println("Successfully export user collcetion at", time.Now())
}

func AdminCollectionLocalBackup(backupDirPath string) {
    collection := database.MongoDatabase.Collection("admin")

    cur, _ := collection.Find(context.TODO(), bson.M{})
    var results []Admin
    for cur.Next(context.TODO()) {
        var elem Admin
        if err := cur.Decode(&elem); err != nil {
            log.Fatal(err)
        }
        
        // set ObjId
        objId := ObjId{Id: elem.TempObjId}
        elem.ObjId = objId

        results = append(results, elem)
    }

    for _, doc := range results {
        b, err := json.Marshal(doc)
        if err != nil {
            log.Fatal(err)
        }
        
        f, err := os.OpenFile(path.Join(backupDirPath, "admin.json"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
        if err != nil {
            log.Fatal(err)
        }
        if _, err := f.Write(b); err != nil {
            log.Fatal(err)
        }
        if _, err := f.Write([]byte("\n")); err != nil {
            log.Fatal(err)
        }
        if err := f.Close(); err != nil {
            log.Fatal(err)
        }
    }

    log.Println("Successfully export admin collcetion at", time.Now())
}

func BusSchedCollectionLocalBackup(backupDirPath string) {
    collection := database.MongoDatabase.Collection("bus_schedule")

    cur, _ := collection.Find(context.TODO(), bson.M{})
    var results []BusSched
    for cur.Next(context.TODO()) {
        var elem BusSched
        if err := cur.Decode(&elem); err != nil {
            log.Fatal(err)
        }
        
        // set ObjId
        objId := ObjId{Id: elem.TempObjId}
        elem.ObjId = objId

        results = append(results, elem)
    }

    for _, doc := range results {
        b, err := json.Marshal(doc)
        if err != nil {
            log.Fatal(err)
        }
        
        f, err := os.OpenFile(path.Join(backupDirPath, "bus_schedule.json"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
        if err != nil {
            log.Fatal(err)
        }
        if _, err := f.Write(b); err != nil {
            log.Fatal(err)
        }
        if _, err := f.Write([]byte("\n")); err != nil {
            log.Fatal(err)
        }
        if err := f.Close(); err != nil {
            log.Fatal(err)
        }
    }

    log.Println("Successfully export bus_schedule collcetion at", time.Now())
}

func AndxCollectionLocalBackup(backupDirPath string) {
    collection := database.MongoDatabase.Collection("anecdote")

    cur, _ := collection.Find(context.TODO(), bson.M{})
    var results []Andx
    for cur.Next(context.TODO()) {
        var elem Andx
        if err := cur.Decode(&elem); err != nil {
            log.Fatal(err)
        }
        
        // set ObjId
        objId := ObjId{Id: elem.TempObjId}
        elem.ObjId = objId

        results = append(results, elem)
    }

    for _, doc := range results {
        b, err := json.Marshal(doc)
        if err != nil {
            log.Fatal(err)
        }
        
        f, err := os.OpenFile(path.Join(backupDirPath, "anecdote.json"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
        if err != nil {
            log.Fatal(err)
        }
        if _, err := f.Write(b); err != nil {
            log.Fatal(err)
        }
        if _, err := f.Write([]byte("\n")); err != nil {
            log.Fatal(err)
        }
        if err := f.Close(); err != nil {
            log.Fatal(err)
        }
    }

    log.Println("Successfully export anecdote collcetion at", time.Now())
}

func FeedbackCollectionLocalBackup(backupDirPath string) {
    collection := database.MongoDatabase.Collection("feedback")

    cur, _ := collection.Find(context.TODO(), bson.M{})
    var results []Feedback
    for cur.Next(context.TODO()) {
        var elem Feedback
        if err := cur.Decode(&elem); err != nil {
            log.Fatal(err)
        }
        
        // set ObjId
        objId := ObjId{Id: elem.TempObjId}
        elem.ObjId = objId

        results = append(results, elem)
    }

    for _, doc := range results {
        b, err := json.Marshal(doc)
        if err != nil {
            log.Fatal(err)
        }
        
        f, err := os.OpenFile(path.Join(backupDirPath, "feedback.json"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
        if err != nil {
            log.Fatal(err)
        }
        if _, err := f.Write(b); err != nil {
            log.Fatal(err)
        }
        if _, err := f.Write([]byte("\n")); err != nil {
            log.Fatal(err)
        }
        if err := f.Close(); err != nil {
            log.Fatal(err)
        }
    }

    log.Println("Successfully export feedback collcetion at", time.Now())
}

func InstanewsCollectionLocalBackup(backupDirPath string) {
    collection := database.MongoDatabase.Collection("instanews")

    cur, _ := collection.Find(context.TODO(), bson.M{})
    var results []News
    for cur.Next(context.TODO()) {
        var elem News
        if err := cur.Decode(&elem); err != nil {
            log.Fatal(err)
        }
        
        // set ObjId
        objId := ObjId{Id: elem.TempObjId}
        elem.ObjId = objId

        results = append(results, elem)
    }

    for _, doc := range results {
        b, err := json.Marshal(doc)
        if err != nil {
            log.Fatal(err)
        }
        
        f, err := os.OpenFile(path.Join(backupDirPath, "instanews.json"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
        if err != nil {
            log.Fatal(err)
        }
        if _, err := f.Write(b); err != nil {
            log.Fatal(err)
        }
        if _, err := f.Write([]byte("\n")); err != nil {
            log.Fatal(err)
        }
        if err := f.Close(); err != nil {
            log.Fatal(err)
        }
    }

    log.Println("Successfully export instanews collcetion at", time.Now())
}

func QaDataCollectionLocalBackup(backupDirPath string) {
    collection := database.MongoDatabase.Collection("qa_data")

    cur, _ := collection.Find(context.TODO(), bson.M{})
    var results []QA
    for cur.Next(context.TODO()) {
        var elem QA
        if err := cur.Decode(&elem); err != nil {
            log.Fatal(err)
        }
        
        // set ObjId
        objId := ObjId{Id: elem.TempObjId}
        elem.ObjId = objId

        results = append(results, elem)
    }

    for _, doc := range results {
        b, err := json.Marshal(doc)
        if err != nil {
            log.Fatal(err)
        }
        
        f, err := os.OpenFile(path.Join(backupDirPath, "qa_data.json"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
        if err != nil {
            log.Fatal(err)
        }
        if _, err := f.Write(b); err != nil {
            log.Fatal(err)
        }
        if _, err := f.Write([]byte("\n")); err != nil {
            log.Fatal(err)
        }
        if err := f.Close(); err != nil {
            log.Fatal(err)
        }
    }

    log.Println("Successfully export qa_data collcetion at", time.Now())
}

func TokenCollectionLocalBackup(backupDirPath string) {
    collection := database.MongoDatabase.Collection("token")

    cur, _ := collection.Find(context.TODO(), bson.M{})
    var results []Token
    for cur.Next(context.TODO()) {
        var elem Token
        if err := cur.Decode(&elem); err != nil {
            log.Fatal(err)
        }
        
        // set ObjId
        objId := ObjId{Id: elem.TempObjId}
        elem.ObjId = objId

        results = append(results, elem)
    }

    for _, doc := range results {
        b, err := json.Marshal(doc)
        if err != nil {
            log.Fatal(err)
        }
        
        f, err := os.OpenFile(path.Join(backupDirPath, "token.json"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
        if err != nil {
            log.Fatal(err)
        }
        if _, err := f.Write(b); err != nil {
            log.Fatal(err)
        }
        if _, err := f.Write([]byte("\n")); err != nil {
            log.Fatal(err)
        }
        if err := f.Close(); err != nil {
            log.Fatal(err)
        }
    }

    log.Println("Successfully export token collcetion at", time.Now())
}