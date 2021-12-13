package pkg

import (
    "net/http"
    "github.com/gin-contrib/sessions"
    "github.com/gin-contrib/sessions/cookie"
    "github.com/gin-gonic/gin"
)

// gin session key
const KEY = "NTHUCHATBOTKEY0405<3"

// 使用 Cookie 保存 Session
func EnableCookieSession() gin.HandlerFunc {
    store := cookie.NewStore([]byte(KEY))
    return sessions.Sessions("nthuChatbotSession", store)
}

// session middleware
func AuthSessionMiddle() gin.HandlerFunc {
    return func(c *gin.Context) {
        session := sessions.Default(c)
        account := session.Get("account")
        
        if account == nil {
            c.JSON(http.StatusUnauthorized, gin.H{
                "result": "",
                "msg": "Unauthorized",
            })
            c.Abort()
            return
        }

        c.Set("account", account)

        c.Next()
        return
    }
}

// 保存 session
func SaveAuthSession(c *gin.Context, account string, department string) {
    session := sessions.Default(c)
    session.Set("account", account)
    session.Set("department", department)
    session.Save()
}

// 清除 session
func ClearAuthSession(c *gin.Context) {
    session := sessions.Default(c)
    session.Clear()
    session.Save()
}

func HasSession(c *gin.Context) bool {
    session := sessions.Default(c)
    if sessionValue := session.Get("account"); sessionValue == nil {
        return false
    }
    return true
}

func GetSessionAccount(c *gin.Context) string {
    session := sessions.Default(c)
    account := session.Get("account")
    if account == nil {
        return "0"
    }
    return account.(string)
}

func GetSessionDeparment(c *gin.Context) string {
    session := sessions.Default(c)
    department := session.Get("department")
    if department == nil {
        return "none"
    }
    return department.(string)
}