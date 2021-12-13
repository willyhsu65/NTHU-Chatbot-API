package v1

import (
    "net/http"
    _"log"
    _"fmt"
    // "strconv"
    "github.com/gin-gonic/gin"
    models "nthu-chatbot-api/models"
)

/*  @desc 新增一則回饋
    @method POST
    @param userID, category, content
*/
func InsertFeedbackApi(c *gin.Context) {
    userID := c.PostForm("userID")
    category := c.PostForm("category")
    content := c.PostForm("content")

    fb := models.Feedback{Category: category, UserID: userID}
    err := fb.InsertOne(content)
    if err == nil {
        c.JSON(http.StatusOK, gin.H{
            "result": "",
            "msg": "Successful insert one feedback",
        })
    } else {
        c.JSON(http.StatusOK, gin.H{
            "result": "",
            "msg": "Insert one feedback fail",
        })
    }
}