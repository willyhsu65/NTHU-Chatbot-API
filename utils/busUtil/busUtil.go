package busUtil

import (
    "fmt"
    "log"
    "time"
)

/* @desc 根據weekday回覆週次類型
   @param weekday int
   @return daily/sat/sun string
*/
func WeekdayMap(weekday int) (wdType string) {
	if weekday==1 || weekday==2 || weekday==3 || weekday==4 || weekday==5 {
        wdType = "daily"
    } else if weekday == 6 {
        wdType = "sat"
    } else if weekday == 7 {
        wdType = "sun" 
    }
    // FIXME: for test
    return "daily"
}


/* @desc 計算經過幾站
   @param route []string
   @param dep string
   @return step int
*/
func CalStep(route []string, dep string) (step int) {
    var depIndex int

    for idx, r := range route {
        if r == dep {
            depIndex = idx
        }
    }

    return depIndex
}


/* @desc 計算發車時間及等待時間
   @param date string 今日日期(2016-12-25)
   @param schedule []string 校車時刻表
   @param duration time.Duration 每站行走時間
   @param step int 經過幾站
   @param busTime string 查詢校車時間(15:04:05)
   @return arriveTime string 抵達時刻
   @return waitTime string 等待時間
   @return err Error 錯誤訊息
*/
func CalArriveTime(date string, schedule []string, duration time.Duration, step int, busTime string) (arriveTime string, waitTime string, err error) {
    // 往前找發車時間
    var firstIdx = 0
    var lastIdx = len(schedule) - 1

    var bTime time.Time  // 現在時間
    var btTime time.Time // bus schedule時間

    for idx, bt := range schedule {
        b_f := fmt.Sprintf("%s %s", date, busTime) // bus time
        bt_f := fmt.Sprintf("%s %s:%s", date, bt, "00") // time range by Schedule
        
        bTime, err = time.Parse("2006-01-02 15:04:05", b_f)
        btTime, err = time.Parse("2006-01-02 15:04:05", bt_f)
        if err != nil {
            log.Println(err.Error())
        }

        var stepDuration time.Duration
        // if bTime在btTime之前 and 第一個時刻
        if btTime.After(bTime) && idx==firstIdx {
            fmt.Println("choose next time")
            stepDuration = time.Duration(step)
            arriveTime = btTime.Add(stepDuration * duration).Format("15:04:05")
            waitTime = (btTime.Sub(bTime) + stepDuration * duration).String()
            return
        }
        
        // if 超過bTime的前一個時刻 or 最後一個時刻
        if btTime.After(bTime) || idx==lastIdx {
            // 前一個時刻+到站duration是否大於或等於bTime
            prevt := schedule[idx-1] // previous time
            prevt_f := fmt.Sprintf("%s %s:%s", date, prevt, "00")
            prevtTime, _ := time.Parse("2006-01-02 15:04:05", prevt_f)

            stepDuration = time.Duration(step)
            arivTime := prevtTime.Add(stepDuration * duration) // 到達時間

            fmt.Println("bTime:", bTime)
            fmt.Println("arivTime:", arivTime)

            fmt.Println("idx:", idx)
            fmt.Println("lastIdx:", lastIdx)
            

            // 超過最後一個班次的時間
            present_time := btTime.Add(stepDuration * duration) // 現在idx的發車時間
            if idx==lastIdx && bTime.After(present_time) {
                fmt.Println("no bus")
                return
            }

            if arivTime.Equal(bTime) || arivTime.After(bTime) {
                fmt.Println("choose prevtTime")
                arriveTime = arivTime.Format("15:04:05")
                waitTime = (arivTime.Sub(bTime)).String()
                fmt.Println("arriveTime:", arriveTime)
                fmt.Println("waitTime:", waitTime)
                return
            }

            if idx == lastIdx {
                fmt.Println("choose next time")
                stepDuration = time.Duration(step)
                arriveTime = btTime.Add(stepDuration * duration).Format("15:04:05")
                waitTime = (btTime.Sub(bTime) + stepDuration * duration).String()
                return
            }

            // 假如還有下一個班次，查看下一個班次
            // if idx != lastIdx {
            //     fmt.Println("choose next time")
            //     stepDuration = time.Duration(step)
            //     arriveTime = btTime.Add(stepDuration * duration).Format("15:04:05")
            //     waitTime = (btTime.Sub(bTime) + stepDuration * duration).String()
            //     return
            // }
        }
    }

    return
}
