package v1

import (
    "net/http"
    "encoding/json"
    "strings"
    "log"
    "fmt"

    "github.com/gin-gonic/gin"

    models "nthu-chatbot-api/models"
)

type m struct {
    Type                string  `json:"type"`
    Text                string  `json:"text"`
    OriginalContentUrl  string  `json:"originalContentUrl"`
    PreviewImageUrl     string  `json:"previewImageUrl"`
    PackageID           string  `json:"packageId"`
    StickerID           string  `json:"stickerId"`
}

type Msg struct {
    To []string `json:"to"`
    Messages []m `json:"messages"`
}

/*  @desc 推送訊息
    @method POST
    @param userID array
    @param msg array
*/
func PushMsgApi(c *gin.Context) {
    var msg Msg
    c.BindJSON(&msg)
    log.Println("Got an push msg call.")

    // 取得line token
    t := models.Token{ID: "0"}
    token, err := t.GetToken()
    if err != nil {
        log.Println(err)

        c.JSON(http.StatusServiceUnavailable, gin.H{
            "result": "",
            "msg": err.Error(),
        })
    }

    // 推送訊息
    linePushMsgAPI := "https://api.line.me/v2/bot/message/multicast"
    var bearer = "Bearer " + token

    // json body
    jsonBody, err := json.Marshal(msg)
    if err != nil {
        log.Println("json body error:", err)
    }

    // new post request
    req, err := http.NewRequest("POST", linePushMsgAPI, strings.NewReader(string(jsonBody)))
    if err != nil {
        log.Println(err)
        c.JSON(http.StatusServiceUnavailable, gin.H{
            "result": "",
            "msg": err.Error(),
        })
    }
    req.Header.Set("Content-Type", "application/json")
    req.Header.Add("Authorization", bearer)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Println("Error on response.\n[ERRO] -", err)
    }

    // get response
    var respResult map[string]interface{}
    json.NewDecoder(resp.Body).Decode(&respResult)

    if (len(respResult) == 0) {
        respMsg := fmt.Sprintf("Successful push msg to %d users", len(msg.To))
        log.Println(respMsg)
        c.JSON(http.StatusOK, gin.H{
            "result": "",
            "msg": respMsg,
        })
    } else {
        log.Println("Push msg to users fail")
        log.Println(respResult)
        c.JSON(http.StatusServiceUnavailable, gin.H{
            "result": "",
            "msg": respResult["message"],
        })
    }
}

