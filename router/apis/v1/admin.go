package v1

import (
    "net/http"
    _"encoding/json"
    _"strings"
    "log"
    _"fmt"

    "github.com/gin-gonic/gin"

    "nthu-chatbot-api/pkg"
    models "nthu-chatbot-api/models"

    "nthu-chatbot-api/pkg/go-drive-backup"
    "nthu-chatbot-api/pkg/go-drive-backup/drive"
    "nthu-chatbot-api/pkg/go-drive-backup/mongo"
)

/*  @desc 啟動備份
    @method GET
*/
func BackupAdminApi(c *gin.Context) {
    goDriveBackup.CreateLocalBackupDir()
    mongo.LocalBackupFunc()
    drive.DriveBackupFunc()

    c.JSON(http.StatusOK, gin.H{
        "result": "",
        "msg": "Successful backup",
    })
}

/*  @desc 取得使用者資料
    @method GET
    @return data [{usreID, time}]
*/
func GetUserInfoAdminApi(c *gin.Context) {
    account := pkg.GetSessionAccount(c)

    admin := models.Admin{Account: account}
    data, err := admin.GetUserInfo()
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

/*  @desc 取得今日新增用戶 & 比昨日新增數量
    @method GET
    @return todayCount int
*/
func GetUserNewCountAdminApi(c *gin.Context) {
    account := pkg.GetSessionAccount(c)

    admin := models.Admin{Account: account}
    todayCount, err := admin.GetUserNewCount()
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
            "result": gin.H{
                "todayCount": todayCount,
            },
            "msg": "Successful get user's info",
        })
    }
}