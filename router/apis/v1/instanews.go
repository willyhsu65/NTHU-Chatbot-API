package v1

import (
    "net/http"
    _"log"
    _"fmt"
    "strconv"

    "github.com/gin-gonic/gin"

    "nthu-chatbot-api/pkg"
    models "nthu-chatbot-api/models"
)

/*  @desc 新增新聞
    @method POST
    @return err
*/
func PostNewsInstanewsApi(c *gin.Context) {
    var news models.News
    c.BindJSON(&news)

    account := pkg.GetSessionAccount(c)

    instanews := models.Instanews{Account: account}
    err := instanews.PostNews(news)
    if err != nil {
        c.JSON(http.StatusServiceUnavailable, gin.H{
            "result": "",
            "msg": err.Error(),
        })
    } else {
        c.JSON(http.StatusOK, gin.H{
            "result": "",
            "msg": "Successful post news",
        })
    }
}

/*  @desc 取得一筆新聞
    @method GET
    @param id
    @return data {ID, PostTime, Title, Date, Time, Location, ImgUrl, Content}
    @return err
*/
func GetOneNewsInstanewsApi(c *gin.Context) {
    id := c.Query("id")
    id64, _ := strconv.ParseInt(id, 10, 64)
    id32 := int32(id64)

    account := pkg.GetSessionAccount(c)

    instanews := models.Instanews{Account: account}
    data, err := instanews.GetOne(id32)
    if err != nil {
        c.JSON(http.StatusServiceUnavailable, gin.H{
            "result": "",
            "msg": err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "result": data,
        "msg": "Successful get one news",
    })
    return
}

/*  @desc 取得所有新聞
    @method GET
    @param category string all/speech/exhibition/activity
    @return data [{ID, PostTime, Title, Date, Time, Location, ImgUrl, Content}]
    @return err
*/
func GetNewsInstanewsApi(c *gin.Context) {
    category := c.Query("category")

    account := pkg.GetSessionAccount(c)

    instanews := models.Instanews{Account: account}
    data, err := instanews.GetNews(category)
    if err != nil {
        c.JSON(http.StatusServiceUnavailable, gin.H{
            "result": "",
            "msg": err.Error(),
        })
    } else {
        c.JSON(http.StatusOK, gin.H{
            "result": data,
            "msg": "Successful get news",
        })
    }
}

/*  @desc 更新一筆新聞
    @method POST
    @param id, postTime, title, category, date, time, location, sourceUrl, imgUrl, content
    @return err
*/
func UpdateOneNewsInstanewsApi(c *gin.Context) {
    var news models.News
    c.BindJSON(&news)

    account := pkg.GetSessionAccount(c)

    instanews := models.Instanews{Account: account}
    err := instanews.UpdateOne(news)
    if err != nil {
        c.JSON(http.StatusServiceUnavailable, gin.H{
            "result": "",
            "msg": err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "result": "",
        "msg": "Successful update one news",
    })
    return
}

/*  @desc 刪除新聞
    @method POST
    @return err
*/
func DeleteNewsInstanewsApi(c *gin.Context) {
    var news models.News
    c.BindJSON(&news)

    account := pkg.GetSessionAccount(c)
    instanews := models.Instanews{Account: account}
    err := instanews.DeleteNews(news.ID)
    if err != nil {
        c.JSON(http.StatusServiceUnavailable, gin.H{
            "result": "",
            "msg": err.Error(),
        })
    } else {
        c.JSON(http.StatusOK, gin.H{
            "result": "",
            "msg": "Successful delete news",
        })
    }
}