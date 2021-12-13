package v1

import (
    "net/http"

    "github.com/gin-gonic/gin"

    models "nthu-chatbot-api/models"
)

/*  @desc 取得所有電話資料
    @method GET
    @param phone_type string 空/academic/administration
    @param data array-object [{name, phone}]
*/
func GetPhonesApi(c *gin.Context) {
    var data []models.Phone
    var err error

    phone_type := c.Query("phone_type")
    phone := models.Phone{}

    if phone_type != "" {
        data, err = phone.GetPhonesByType(phone_type)
    } else {
        data, err = phone.GetPhones()
    }

    if err != nil {
        c.JSON(http.StatusServiceUnavailable, gin.H{
            "result": "",
            "msg": err.Error(),
        })
    } else {
        c.JSON(http.StatusOK, gin.H{
            "result": data,
            "msg": "Successful get phones",
        })
    }
}