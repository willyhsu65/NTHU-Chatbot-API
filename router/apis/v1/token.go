package v1

import (
    "net/http"
    "log"

    "github.com/gin-gonic/gin"

    models "nthu-chatbot-api/models"
)

/*  @desc 取得 LINE token & webhook
    @method GET
    @param mode string
*/
func GetLineAuthTokenApi(c *gin.Context) {
    mode := c.Query("mode")

    var data models.Token
    var err error
    if mode == "official" {
        t := models.Token{ID: "0"}
        data, err = t.GetTokenAndWebhook()
        if err != nil {
            log.Println(err)
            c.JSON(http.StatusServiceUnavailable, gin.H{
                "result": "",
                "msg": err.Error(),
            })
        }
    } else if mode == "beta" {
        t := models.Token{ID: "1"}
        data, err = t.GetTokenAndWebhook()
        if err != nil {
            log.Println(err)
            c.JSON(http.StatusServiceUnavailable, gin.H{
                "result": "",
                "msg": err.Error(),
            })
        }
    }

    c.JSON(http.StatusOK, gin.H{
        "result": gin.H{
            "token": data.Token,
            "webhook": data.Webhook,
        },
        "msg": "Successfully get LINE Token & Webhook",
    })
}

