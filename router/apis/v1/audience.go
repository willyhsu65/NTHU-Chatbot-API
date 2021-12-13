package v1

import (
    "net/http"
    "log"
    _"fmt"
    
    "github.com/gin-gonic/gin"
    models "nthu-chatbot-api/models"
)

/*  @desc 建立受眾
    @method POST
    @param description string
    @param audiences [].id string
    @return audienceGroupId int
*/
func CreateAudiApi(c *gin.Context) {
    var audi models.CreateAudiBody
    c.BindJSON(&audi)

    // check is the description already exist
    _, audienceGroupId, err := audi.CheckAudienceByDesc(audi.Description)
    if err != nil {
        c.JSON(http.StatusServiceUnavailable, gin.H{
            "result": err.Error(),
            "msg": "Create Audi failed",
        })
        return
    }

    audienceGroupId, err = audi.CreateAudience(audi.Description, audi.Audiences)
    if err != nil {
        c.JSON(http.StatusServiceUnavailable, gin.H{
            "result": err.Error(),
            "msg": "Create Audi failed",
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "result": gin.H{
            "audienceGroupId": audienceGroupId,
        },
        "msg": "Successful create Audi",
    })
    return
}


/*  @desc 刪除受眾
    @method POST
    @param audienceGroupId int
*/
func DeleteAudiApi(c *gin.Context) {
    var audi models.UpdateAudiBody
    c.BindJSON(&audi)

    // find id descrition name
    var audiBody models.CreateAudiBody
    _, description, err := audiBody.CheckAudienceById(audi.GroupID)
    if err != nil {
        c.JSON(http.StatusServiceUnavailable, gin.H{
            "result": err.Error(),
            "msg": "Delete Audi failed",
        })
        return
    }

    // delete audience from line
    err = audi.DeleteAudience(audi.GroupID)
    if err != nil {
        c.JSON(http.StatusServiceUnavailable, gin.H{
            "result": err.Error(),
            "msg": "Delete Audi failed",
        })
        return
    }

    // delete audi tag from user's tag category
    var user models.User
    resultCount, err := user.RemoveTagByAudiDesc(description)
    if err != nil {
        c.JSON(http.StatusServiceUnavailable, gin.H{
            "result": err.Error(),
            "msg": "Delete Audi failed",
        })
        return
    } else {
        log.Println("Successfully delete tag from", resultCount, "users")
    }
    
    c.JSON(http.StatusOK, gin.H{
        "result": "",
        "msg": "Successful delete Audience",
    })
    return
}

/*
* @desc lineID綁定受眾
* @method POST
* @param userID string
* @param description string
* @param audiences [].id string
* @return audienceGroupId int
*/
type BindUserAudiBody struct {
    UserID  string  `json:"userID" form:"userID"`
    *models.CreateAudiBody
}

func BindUserAudiApi(c *gin.Context) {
    var bindUserAudiBody BindUserAudiBody
    c.BindJSON(&bindUserAudiBody)

    var createAudi models.CreateAudiBody

    // 查看是已存在此 受眾description
    isExist, audienceGroupId, err := createAudi.CheckAudienceByDesc(bindUserAudiBody.Description)
    if err != nil {
        c.JSON(http.StatusServiceUnavailable, gin.H{
            "result": err.Error(),
            "msg": "Create Audi failed",
        })
        return
    }

    // 若無：建立受眾並加入lineID
    if isExist == false {
        audienceGroupId, err = createAudi.CreateAudience(bindUserAudiBody.Description, bindUserAudiBody.Audiences)
        if err != nil {
            c.JSON(http.StatusServiceUnavailable, gin.H{
                "result": err.Error(),
                "msg": "Create Audi failed",
            })
            return
        }
    } else { // 若有：直接加入到 受眾
        var updateAudi models.UpdateAudiBody
        err = updateAudi.UpdateAudi(audienceGroupId, bindUserAudiBody.Audiences)
        if err != nil {
            c.JSON(http.StatusServiceUnavailable, gin.H{
                "result": err.Error(),
                "msg": "Create Audi failed",
            })
            return
        }
    }

    // 新增到 user's tag category
    var user models.User
    err = user.UpdateTag(bindUserAudiBody.UserID, bindUserAudiBody.Description)
    if err != nil {
        c.JSON(http.StatusServiceUnavailable, gin.H{
            "result": err.Error(),
            "msg": "Create Audi failed",
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "result": gin.H{
            "audienceGroupId": audienceGroupId,
        },
        "msg": "Successful bind user's audience",
    })
    return
}

/*
* @desc 取得所有audience資料
* @method GET
* @return data [{audienceGroupId, description, audienceCount}]
*/
func GetAllAudiApi(c *gin.Context) {
    var audi models.Audience

    data, err := audi.GetAllAudience()
    if err != nil {
        c.JSON(http.StatusServiceUnavailable, gin.H{
            "result": err.Error(),
            "msg": "Get all Audience data failed",
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "result": data,
        "msg": "Successful get Audience data",
    })
    return
}