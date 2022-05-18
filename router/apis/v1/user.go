package v1

import (
    "net/http"
    _"log"
    _"fmt"
    "strconv"

    "github.com/gin-gonic/gin"

    models "nthu-chatbot-api/models"
)

/*  @desc 新增user
    @method POST
    @param userID
*/
func InsertOneUserApi(c *gin.Context) {
    userID := c.PostForm("userID")
    user := models.User{UserID: userID}

    err := user.InsertOne()
    if err != nil {
        c.JSON(http.StatusServiceUnavailable, gin.H{
            "result": "",
            "msg": err.Error(),
        })
    } else {
        c.JSON(http.StatusOK, gin.H{
            "result": "",
            "msg": "Successful insert one user",
        })
    }
}


/*  @desc 此user是否存在
    @method GET
    @param userID(string)
    @return isExist(bool)
*/
func IsExistUserApi(c *gin.Context) {
    var userID string

    userID = c.Query("userID")
    user := models.User{UserID: userID}

    isExist, err := user.IsExist()
    if err != nil {
        c.JSON(http.StatusServiceUnavailable, gin.H{
            "result": "",
            "msg": err.Error(),
        })
    } else {
        c.JSON(http.StatusOK, gin.H{
            "result": gin.H{
                "isExist": isExist,
            },
            "msg": "Successful search user whether exist",
        })
    }
}


/*  @desc 取得user的flag
    @method GET
    @param userID(string)
    @return flag(srting)
*/
func GetFlagUserApi(c *gin.Context) {
    var userID string

    userID = c.Query("userID")
    user := models.User{UserID: userID}

    flag, err := user.GetFlag()
    if err != nil {
        c.JSON(http.StatusServiceUnavailable, gin.H{
            "result": "",
            "msg": err.Error(),
        })
    } else {
        c.JSON(http.StatusOK, gin.H{
            "result": gin.H{
                "flag": flag,
            },
            "msg": "Successful get user's flag",
        })
    }
}


/*  @desc 初始化user的flag
    @method POST
    @param userID
*/
func InitFlagUserApi(c *gin.Context) {
    var userID string

    userID = c.PostForm("userID")
    user := models.User{UserID: userID}

    err := user.InitFlag()
    if err != nil {
        c.JSON(http.StatusServiceUnavailable, gin.H{
            "result": "",
            "msg": err.Error(),
        })
    } else {
        c.JSON(http.StatusOK, gin.H{
            "result": "",
            "msg": "Successful init user's flag",
        })
    }
}


/*  @desc 設置user的flag
    @method POST
    @param userID, flag
*/
func SetFlagUserApi(c *gin.Context) {
    var userID string
    var flag string

    userID = c.PostForm("userID")
    flag = c.PostForm("flag")
    
    user := models.User{UserID: userID}

    err := user.SetFlag(flag)
    if err != nil {
        c.JSON(http.StatusServiceUnavailable, gin.H{
            "result": "",
            "msg": err.Error(),
        })
    } else {
        c.JSON(http.StatusOK, gin.H{
            "result": "",
            "msg": "Successful set user's flag",
        })
    }
}


/*  @desc 取得使用者資料
    @method GET
    @param userID (string)
    @return name (string)
    @return time (string)
*/
func GetInfoUserApi(c *gin.Context) {
    userID := c.PostForm("userID")

    user := models.User{UserID: userID}
    time, err := user.GetInfo()
    if err != nil {
        c.JSON(http.StatusServiceUnavailable, gin.H{
            "result": "",
            "msg": err.Error(),
        })
    } else {
        c.JSON(http.StatusOK, gin.H{
            "result": gin.H{
                "time": time,
            },
            "msg": "Successful get user's info",
        })
    }
}

/*  @desc 新增一筆查詢校園地圖資料
    @method POST
    @param userID string
    @param location string
*/
func InsertMapRecordUserApi(c *gin.Context) {
    userID := c.PostForm("userID")
    location := c.PostForm("location")
    
    user := models.User{UserID: userID}
    err := user.InsertMapRecord(location)
    if err != nil {
        c.JSON(http.StatusServiceUnavailable, gin.H{
            "result": "",
            "msg": err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "result": "",
        "msg": "Successful insert user's map record",
    })
}


/*  @desc 取得此使用者查詢校園地圖紀錄
    @method GET
    @param location array 校園地點名稱
*/
func GetMapRecordUserApi(c *gin.Context) {
    userID := c.Query("userID")
    user := models.User{UserID: userID}

    locations, err := user.GetMapRecord()
    if err != nil {
        c.JSON(http.StatusServiceUnavailable, gin.H{
            "result": "",
            "msg": err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "result": locations,
        "msg": "Successful get user's map records",
    })
}


/*  @desc 綁定學生學號
    @method POST
    @param userID string
    @param studentID string
*/
func BindStudentIDUserApi(c *gin.Context) {
    userID := c.PostForm("userID")
    studentID := c.PostForm("studentID")
    
    user := models.User{UserID: userID}
    err := user.BindStudentID(studentID)
    if err != nil {
        c.JSON(http.StatusServiceUnavailable, gin.H{
            "result": "",
            "msg": err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "result": "",
        "msg": "Successful bind user's studentID",
    })
}

/*  @desc 設定使用者的上次被推播的時間戳
    @method POST
    @param userID string
    @param tag int
*/
func UpdateBroadcastTag(c *gin.Context) {
    userID := c.PostForm("userID")
    tag := c.PostForm("tag")
    
    user := models.User{UserID: userID}
    tag_int, err := strconv.Atoi(tag)
    if err == nil {
        // Pass down the error to next stage.
        err = user.UpdateBroadcastTag(tag_int)
    }

    if err != nil {
        c.JSON(http.StatusServiceUnavailable, gin.H{
            "result": "",
            "msg": err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "result": "",
        "msg": "Successful updated user's tag",
    })
}

/*  @desc 取得要被推播的使用者 id
    @method GET
*/
func GetBroadcastAudienceIds(c *gin.Context) {   
    userID := c.Query("userID")

    user := models.User{}
    err, audienceIds := user.GetBroadcastAudienceIds(userID)
    if err != nil {
        c.JSON(http.StatusServiceUnavailable, gin.H{
            "result": "",
            "msg": err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "result": audienceIds,
        "msg": "Successful selected id of unbroadcasted users",
    })
}
