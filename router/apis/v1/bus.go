package v1

import (
	"net/http"
    "log"
    "fmt"
    // "strconv"
    "github.com/gin-gonic/gin"
    models "nthu-chatbot-api/models"
)


/*  @method GET
    @param time 欲查詢的時刻
    @param busType 校車類型(campus)
    @param direction 上山或下山(climb/descend)
    @param line 紅線或綠線(red/green)
    @param dep 出發站名(北校門口, 綜二館, 楓林小徑, 奕園停車場, 南門停車場, 人社院, 台積館)
    @return arriveTime 抵達時間
    @return waitTime 等待時間
*/
func GetScheduleBusApi(c *gin.Context) {
    // params
    var time string
    var busType string
    var direction string
    var line string
    var dep string

    time = c.Query("time")
    busType = c.Query("busType")
    direction = c.Query("direction")
    line = c.Query("line")
    dep = c.Query("dep")

    bus := models.Bus{Time: time, BusType: busType, Direction: direction, Line: line, Dep: dep}

    var msg string
    arriveTime, waitTime, err := bus.GetDepTime()
    if err != nil {
        log.Println(err)
        msg = err.Error()
    } else {
        msg = fmt.Sprintf("Successful get bus schedule")
    }

    
    c.JSON(http.StatusOK, gin.H{
        "result": gin.H{
            "arriveTime": arriveTime,
            "waitTime": waitTime,
        },
        "msg": msg,
    })
}