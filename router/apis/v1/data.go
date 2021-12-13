package v1

import (
    "net/http"
    _"log"
    _"fmt"

    "github.com/gin-gonic/gin"

    models "nthu-chatbot-api/models"
)

/*  @desc 取得QA資料
    @method GET
    @param category string
    @return data [{ques, ans}]
*/
func GetQADataApi(c *gin.Context) {
    category := c.Query("category")

    qa := models.QA{}
    data, err := qa.GetOneData(category)
    if err != nil {
        c.JSON(http.StatusServiceUnavailable, gin.H{
            "result": "",
            "msg": err.Error(),
        })
    } else {
        c.JSON(http.StatusOK, gin.H{
            "result": data,
            "msg": "Successful get QA data",
        })
    }
}

/*  @desc 取得防疫資訊
    @method GET
    @param category string
    @return data []string
*/
func GetEpidDataApi(c *gin.Context) {
    category := c.Query("category")

    epid := models.Epidemic{}
    data, err := epid.GetOne(category)
    if err != nil {
        c.JSON(http.StatusServiceUnavailable, gin.H{
            "result": "",
            "msg": err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "result": data,
        "msg": "Successful get Epidemic data",
    })
    return
}

/*  @desc 推薦學生新聞
    @method GET
    @return [{title, category, date, time, location, imgUrl, content}]
*/
func GetRecNewsDataApi(c *gin.Context) {
    category := c.Query("category")

    instanews := models.Instanews{}

    data, err := instanews.GetRecNews(category)
    if err != nil {
        c.JSON(http.StatusServiceUnavailable, gin.H{
            "result": "",
            "msg": err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "result": data,
        "msg": "Successful get recommend news data",
    })
    return
}