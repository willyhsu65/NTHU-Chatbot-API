package models

import (
    "context"
    "fmt"
    "strings"
    "log"
    "time"
    _"reflect"
    _"strconv"
    
    "go.mongodb.org/mongo-driver/bson"
    _"go.mongodb.org/mongo-driver/bson/primitive"
    // "go.mongodb.org/mongo-driver/mongo"
    // "go.mongodb.org/mongo-driver/mongo/options"

    db "nthu-chatbot-api/database"
    "nthu-chatbot-api/utils/busUtil"
)

type Bus struct {
    Time       string  `json:"time" form:"time"`
    BusType    string  `json:"busType" form:"busType"`
    Direction  string  `json:"direction" form:"direction"`
    Line       string  `json:"line" form:"line"`
    Dep        string  `json:"dep" form:"dep"`
}

type BusSched struct {
    BusType    string    `json:"busType" form:"busType"`
    WdType     string    `json:"wdType" form:"wdType"`
    Direction  string    `json:"direction" form:"direction"`
    Line       string    `json:"line" form:"line"`
    Route      []string  `json:"route" form:"route"`
    Schedule   []string  `json:"schedule" form:"schedule"`
}

/* @desc 取得Time與校車出發最相近的時間（在Time之後）
   @return arriveTime(string) 抵達時間
   @return waitTime(string) 等待時間
   @return err(error)
*/
func (b *Bus) GetDepTime() (arriveTime string, waitTime string, err error) {
    t := time.Now()
    // date := fmt.Sprintf("%d-%d-%d", t.Year(), t.Month(), t.Day())
    date := t.Format("2006-01-02") 

    // 解析今天是週幾
    weekday := int(t.Weekday())
    wdType := busUtil.WeekdayMap(weekday)

    // 查看hour是否有在發車時段內
    timeSplit := strings.Split(b.Time, ":") // 處理傳來的Time: 分離hour & min
    hour := timeSplit[0]

    var Peek []string = make([]string, 1)
    flag := false
    if wdType == "daily" {
        Peek = []string{"07", "08", "09", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22"}
        for _, peek := range Peek {
            if hour == peek {
                flag = true
            }
        }
    }

    if flag == false {
        err = fmt.Errorf("此時段沒有發車")
        return
    }

    // 從 mongo 取得 schedule 資料
    bus_collect := db.MongoDatabase.Collection("bus_schedule")
    
    // 暑假
    // bus_collect := db.MongoDatabase.Collection("bus_schedule_summer")

    var filter bson.M
    var result BusSched
    if b.Direction == "climb" {
        if b.Dep=="綜二館" &&  b.Line=="green"{
            filter = bson.M{"busType": b.BusType, "wdType": wdType, "direction": b.Direction, "line": b.Line, "depStop": "綜二館"}
        } else {
            filter = bson.M{"busType": b.BusType, "wdType": wdType, "direction": b.Direction, "line": b.Line, "depStop": "北校門口"}
        }
    } else if b.Direction == "descend" {
        filter = bson.M{"busType": b.BusType, "wdType": wdType, "direction": b.Direction, "line": b.Line, "depStop": "台積館"}
    }

    ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
    err = bus_collect.FindOne(ctx, filter).Decode(&result)
    if err != nil {
        log.Println("Error on Finding documents", err)
    }
    fmt.Println(result)


    /* 頭站出發: 北校門口、綜二館、台積館  start */
    if b.Dep=="北校門口" || (b.Dep=="綜二館"&&b.Line=="green") || b.Dep=="台積館" {
        // 找是否有剛好的時間
        for _, bt := range result.Schedule {
            b_f := fmt.Sprintf("%s %s", date, b.Time)
            bt_f := fmt.Sprintf("%s %s:%s", date, bt, "00")
            
            bTime, e := time.Parse("2006-01-02 15:04:05", b_f)
            btTime, e := time.Parse("2006-01-02 15:04:05", bt_f)
            if e != nil {
                log.Println(e.Error())
            }

            if bTime.Equal(btTime) {
                arriveTime = fmt.Sprintf("%d:%d:%d", btTime.Hour(), btTime.Minute(), btTime.Second())
                waitTime = "0s"
                return
            }
        }

        // 北校門口 跟 台積館 沒有剛好時間, 往後找。綜二館到後面處理
        if b.Dep=="北校門口" || b.Dep=="台積館" {
            for _, bt := range result.Schedule {
                b_f := fmt.Sprintf("%s %s", date, b.Time)
                bt_f := fmt.Sprintf("%s %s:%s", date, bt, "00")
                
                bTime, e := time.Parse("2006-01-02 15:04:05", b_f)
                btTime, e := time.Parse("2006-01-02 15:04:05", bt_f)
                if e != nil {
                    log.Println(e.Error())
                }
    
                if bTime.Before(btTime) {
                    arriveTime = fmt.Sprintf("%d:%d:%d", btTime.Hour(), btTime.Minute(), btTime.Second())
                    waitTime = btTime.Sub(bTime).String()
                    return
                }
            }
        }
    }
    /* 頭站出發: 北校門口、綜二館、台積館 end */
    

    /* 不是頭站出發, 需要去計算step*行走時間 start */
    var step = busUtil.CalStep(result.Route, b.Dep)
    var duration time.Duration
    
    if b.Direction=="climb" && b.Line=="red" {
        duration, _ = time.ParseDuration("1m45s") // 校車每走一格行進時間
    } else if b.Direction=="climb" && b.Line=="green" {
        duration, _ = time.ParseDuration("1m10s")
    } else if b.Direction=="descend" && b.Line=="red" {
        duration, _ = time.ParseDuration("1m10s")
    } else if b.Direction=="descend" && b.Line=="green" {
        duration, _ = time.ParseDuration("1m45s")
    }

    fmt.Println("step:", step)
    fmt.Println("duration:", duration)

    // 特別處理climb green line
    // 比較"綜二出發"與"北校門口"出發，看誰比較快來
    if b.Direction=="climb" && b.Line=="green" {
        // 北校門口出發
        maingate2g2Step := 1 // 為了計算北校門出發到綜二館
        filter = bson.M{"busType": b.BusType, "wdType": wdType, "direction": b.Direction, "line": b.Line, "depStop": "北校門口"}
        err = bus_collect.FindOne(ctx, filter).Decode(&result)
        fmt.Println(result)
        mainArriveTime, mainWaitTime, err := busUtil.CalArriveTime(date, result.Schedule, duration, maingate2g2Step, b.Time)
        if err != nil {
            log.Println(err)
        }

        // 綜二出發
        filter = bson.M{"busType": b.BusType, "wdType": wdType, "direction": b.Direction, "line": b.Line, "depStop": "綜二館"}
        err = bus_collect.FindOne(ctx, filter).Decode(&result)
        fmt.Println(result)
        g2ArriveTime, g2WaitTime, err := busUtil.CalArriveTime(date, result.Schedule, duration, step, b.Time)
        if err != nil {
            log.Println(err)
        }

        // 比較兩個等待時間
        mainArriveTimeParse, _ := time.Parse("15:04:05", mainArriveTime)
        g2ArriveTimeParse, _ := time.Parse("15:04:05", g2ArriveTime)

        if mainArriveTimeParse.Before(g2ArriveTimeParse) {
            fmt.Println("北校門口較快")
            arriveTime = mainArriveTime
            waitTime = mainWaitTime
        } else {
            fmt.Println("綜二館較快")
            arriveTime = g2ArriveTime
            waitTime = g2WaitTime
        }
    } else {
        fmt.Println("其他")

        arriveTime, waitTime, err = busUtil.CalArriveTime(date, result.Schedule, duration, step, b.Time)
        if err != nil {
            log.Println(err)
        }
    }
    /* 不是頭站出發, 需要去計算step*行走時間 end */

    return
}