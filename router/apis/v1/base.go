package v1

import (
    "net/http"
    "github.com/gin-gonic/gin"
    _"fmt"

    "nthu-chatbot-api/pkg"
    models "nthu-chatbot-api/models"
)

func PingApi(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "result": "",
        "msg": "NTHU Chatbot is ALIVE",
    })
}

func LoginApi(c *gin.Context) {
    var adminUser models.Admin
    c.BindJSON(&adminUser)

    account := adminUser.Account
    password := adminUser.Password

    // if hasSession := pkg.HasSession(c); hasSession == true {
    //     c.String(200, "You are already logged in")
    //     return
    // }
    
    admin := models.Admin{Account: account, Password: password}
    result, isExist, err := admin.UserDetail()
    if err != nil {
        c.JSON(http.StatusServiceUnavailable, gin.H{
            "result": "",
            "msg": err.Error(),
        })
    }

    if isExist == false {
        c.JSON(http.StatusUnauthorized, gin.H{
            "result": "",
            "msg": "無此帳號密碼",
        })
    } else {
        pkg.SaveAuthSession(c, result.Account, result.Department)

        c.JSON(http.StatusOK, gin.H{
            "result": gin.H{
                "department": result.Department,
                "name": result.Name,
            },
            "msg": "登入成功",
        })
    }
}


func LogoutApi(c *gin.Context) {
    if hasSession := pkg.HasSession(c); hasSession == false {
        c.JSON(http.StatusUnauthorized, gin.H{
            "result": "",
            "msg": "用戶未登入",
        })
        return
    }
    pkg.ClearAuthSession(c)

    c.JSON(http.StatusOK, gin.H{
        "result": "",
        "msg": "登出成功",
    })
}


func CheckLoginApi(c *gin.Context) {
    if hasSession := pkg.HasSession(c); hasSession == false {
        c.JSON(http.StatusUnauthorized, gin.H{
            "result": "",
            "msg": "用戶未登入",
        })
        return
    }

    department := pkg.GetSessionDeparment(c);
    c.JSON(http.StatusOK, gin.H{
        "result": gin.H{
            "department": department,
        },
        "msg": "用戶登入中",
    })
}