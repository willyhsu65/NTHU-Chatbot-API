package beta

import (
    "net/http"
    _"encoding/json"
    _"strings"
    "log"
    _"fmt"

    "github.com/gin-gonic/gin"

    "nthu-chatbot-api/pkg"
    models "nthu-chatbot-api/models"
)

/*  @desc 取得測試者使用者資料
    @method GET
    @return data [{usreID, time}]
*/
func GetUserInfoAdminApi(c *gin.Context) {
    account := pkg.GetSessionAccount(c)

    admin := models.Admin{Account: account}
    data, err := admin.GetBetaUserInfo()
    if err != nil { 
        log.Println(err) 
    }

    if err != nil {
        c.JSON(http.StatusServiceUnavailable, gin.H{
            "result": "",
            "msg": err.Error(),
        })
    } else {
        c.JSON(http.StatusOK, gin.H{
            "result": data,
            "msg": "Successful get user's info",
        })
    }
}