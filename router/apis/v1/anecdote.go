package v1

import (
	"net/http"
    _"log"
    _"fmt"
    // "strconv"
    "github.com/gin-gonic/gin"
    models "nthu-chatbot-api/models"
)

func InsertANDXApi(c *gin.Context) {
    userID := c.PostForm("userID")
    content := c.PostForm("content")

    andx := models.Andx{UserID: userID}
    err := andx.InsertOne(content)
    if err == nil {
        c.JSON(http.StatusOK, gin.H{
            "result": "",
            "msg": "Successful insert one anecdote",
        })
    } else {
        c.JSON(http.StatusOK, gin.H{
            "result": "",
            "msg": "Insert one anecdote fail",
        })
    }
}

func GetRandANDXApi(c *gin.Context) {
    andx := models.Andx{}
    content, err := andx.GetOne()

    if err == nil {
        c.JSON(http.StatusOK, gin.H{
            "result": gin.H{
                "content": content,
            },
            "msg": "Successful get one random anecdote",
        })
    } else {
        c.JSON(http.StatusServiceUnavailable, gin.H{
            "result": "",
            "msg": "get random anecdote fail",
        })
    }
    
}