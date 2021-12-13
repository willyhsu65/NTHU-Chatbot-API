package v1

import (
    "net/http"
    _"log"
    _"fmt"

    "github.com/gin-gonic/gin"

    "nthu-chatbot-api/pkg"
    models "nthu-chatbot-api/models"
)

/* anecdote */
/*  @desc 取得所有趣聞
    @method GET
    @return data [{id, userID, time, content}]
    @return err error
*/
func GetAndxGroupApi(c *gin.Context) {
    account := pkg.GetSessionAccount(c)
    group := models.Group{Account: account}

    data, err := group.GetAndx()
    if err != nil {
        c.JSON(http.StatusServiceUnavailable, gin.H{
            "result": "",
            "msg": err.Error(),
        })
    } else {
        c.JSON(http.StatusOK, gin.H{
            "result": data,
            "msg": "Successful get anecdote",
        })
    }
}

/*  @desc 刪除一則趣聞
    @method DELETE
    @return err
*/
func DeleteAndxGroupApi(c *gin.Context) {
    var andx models.Andx
    c.BindJSON(&andx)

    account := pkg.GetSessionAccount(c)
    group := models.Group{Account: account}

    err := group.DeleteAndx(andx.ID)
    if err != nil {
        c.JSON(http.StatusServiceUnavailable, gin.H{
            "result": "",
            "msg": err.Error(),
        })
    } else {
        c.JSON(http.StatusOK, gin.H{
            "result": "",
            "msg": "Successful delete anecdote",
        })
    }
}

/* feedback */
/*  @desc 取得所有回饋
    @method GET
    @return data [{id, userID, time, content}]
    @return err error
*/
func GetFeedbackGroupApi(c *gin.Context) {
    var category string
    category = c.Query("category")

    account := pkg.GetSessionAccount(c)
    group := models.Group{Account: account}

    data, err := group.GetFeedback(category)
    if err != nil {
        c.JSON(http.StatusServiceUnavailable, gin.H{
            "result": "",
            "msg": err.Error(),
        })
    } else {
        c.JSON(http.StatusOK, gin.H{
            "result": data,
            "msg": "Successful get feedback",
        })
    }
}

/*  @desc 刪除一則回饋
    @method POST
    @return err
*/
func DeleteFeedbackGroupApi(c *gin.Context) {
    var feedback models.Feedback
    c.BindJSON(&feedback)

    account := pkg.GetSessionAccount(c)
    group := models.Group{Account: account}

    err := group.DeleteFeedback(feedback.ID)
    if err != nil {
        c.JSON(http.StatusServiceUnavailable, gin.H{
            "result": "",
            "msg": err.Error(),
        })
    } else {
        c.JSON(http.StatusOK, gin.H{
            "result": "",
            "msg": "Successful delete feedback",
        })
    }
}

/*  @desc 新增QA類別
    @method POST
    @param category string 
*/
func AddQACategoryGroupApi(c *gin.Context) {
    var qa models.QA
    c.BindJSON(&qa)

    err := qa.AddCategory(qa.Category)
    if err != nil {
        c.JSON(http.StatusServiceUnavailable, gin.H{
            "result": "",
            "msg": err.Error(),
        })
    } else {
        c.JSON(http.StatusOK, gin.H{
            "result": "",
            "msg": "Successful add QA category",
        })
    }
}

/*  @desc 新增QA資料
    @method POST
    @param category string
    @param ques string
    @param ans string
*/
func AddOneQAGroupApi(c *gin.Context) {
    var qaInfo models.QAInfo
    c.BindJSON(&qaInfo)
    
    qa := models.QA{Category: qaInfo.Category}
    qaData := models.QAData{Ques: qaInfo.Ques, Ans: qaInfo.Ans}

    err := qa.AddOne(qaData)
    if err != nil {
        c.JSON(http.StatusServiceUnavailable, gin.H{
            "result": "",
            "msg": err.Error(),
        })
    } else {
        c.JSON(http.StatusOK, gin.H{
            "result": "",
            "msg": "Successful add QA data",
        })
    }
}

/*  @desc 取得所有QA資料
    @method GET
    @return data [{id, category, data:[{ques, ans}]}]
*/
func GetAllQAGroupApi(c *gin.Context) {
    qa := models.QA{}

    data, err := qa.GetAllData()
    if err != nil {
        c.JSON(http.StatusServiceUnavailable, gin.H{
            "result": "",
            "msg": err.Error(),
        })
    } else {
        c.JSON(http.StatusOK, gin.H{
            "result": data,
            "msg": "Successful add QA data",
        })
    }

}

/*  @desc 刪除QA資料
    @method POST
    @param category string
    @param ques string
*/
func DelOneQAGroupApi(c *gin.Context) {
    var qaInfo models.QAInfo
    c.BindJSON(&qaInfo)

    qa := models.QA{}
    err := qa.DelOneQA(qaInfo.Category, qaInfo.Ques)
    if err != nil {
        c.JSON(http.StatusServiceUnavailable, gin.H{
            "result": "",
            "msg": err.Error(),
        })
    }

    c.JSON(http.StatusOK, gin.H{
        "result": "",
        "msg": "Successful del one QA",
    })
}

/*  @desc 取得所有防疫訊息
    @method GET
    @return [{id, category, title, []content}]
*/
func GetAllEpidGroupApi(c *gin.Context) {
    epid := models.Epidemic{}
    data, err := epid.GetAll()
    if err != nil {
        c.JSON(http.StatusServiceUnavailable, gin.H{
            "result": "",
            "msg": err.Error(),
        })
    }

    c.JSON(http.StatusOK, gin.H{
        "result": data,
        "msg": "Successful get all epidemic data",
    })
}

/*  @desc 更新一筆防疫訊息
    @method POST
    @param category
    @param content []string
*/
func UpdateOneEpidGroupApi(c *gin.Context) {
    var epid models.Epidemic
    c.BindJSON(&epid)

    err := epid.UpdateOne()
    if err != nil {
        c.JSON(http.StatusServiceUnavailable, gin.H{
            "result": "",
            "msg": err.Error(),
        })
    }

    c.JSON(http.StatusOK, gin.H{
        "result": "",
        "msg": "Successful update one epidemic data",
    })
}