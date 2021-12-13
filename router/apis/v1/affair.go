package v1

import (
    "encoding/json"
    "net/http"
    _"io/ioutil"
    "log"
    "fmt"
    "github.com/gin-gonic/gin"
)

/*  @desc 詢問校務資訊
    @method POST
    @param userID
*/
func FindQuesAffairApi(c *gin.Context) {
    affairAPI := "http://flask:5000/api/flask/affair/find_ques"
    ques := c.Query("ques")
    
    // 建立 req
    req, err := http.NewRequest("GET", affairAPI, nil)
    if err != nil {
        log.Println(err)
    }

    q := req.URL.Query()
    q.Add("ques", ques)
    req.URL.RawQuery = q.Encode()

    // 發送 req
    var resp *http.Response
    resp, err = http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        errMsg := fmt.Sprintf("Flask Affair API error: %s", resp.StatusCode)

        c.JSON(http.StatusServiceUnavailable, gin.H{
            "result": "",
            "msg": errMsg,
        })
    } else { // if status code is OK
        var result map[string]interface{}
        json.NewDecoder(resp.Body).Decode(&result)
        
        if err != nil {
            c.JSON(http.StatusServiceUnavailable, gin.H{
                "result": "",
                "msg": err.Error(),
            })
        } else {
            c.JSON(http.StatusOK, gin.H{
                "result": gin.H{
                    "ans": result["result"],
                },
                "msg": "Successful get school affair",
            })
        }
    }
}