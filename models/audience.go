package models

import (
    "net/http"
    "encoding/json"
    "strings"
    "log"
    "fmt"
    "errors"
    
    _"go.mongodb.org/mongo-driver/bson"

    // db "nthu-chatbot-api/database"
)

type AudiencesID struct {
    ID string `json:"id"`
}

type CreateAudiBody struct {
    Description    string         `json:"description"`
    IsIfaAudience  bool           `json:"isIfaAudience"`
    Audiences      []AudiencesID  `json:"audiences"`
}

// for update or delete audience
type UpdateAudiBody struct {
    GroupID    int            `json:"audienceGroupId"`
    Audiences  []AudiencesID  `json:"audiences"`
}

type Audience struct {
    GroupID     int    `json:"audienceGroupId"`
    Description string `json:"description"`
    Count       int    `json:"audienceCount"`
}

// 檢查是否有此 audienceGroupId 的 audience
func (audi *CreateAudiBody) CheckAudienceById(audienceGroupId int) (isExist bool, description string, err error) {
    // 取得line token
    t := Token{ID: "0"}
    token, _ := t.GetToken()

    // 取得 audience 
    audienceGroupIdString := fmt.Sprint(audienceGroupId)
    getAudiAPI := "https://api.line.me/v2/bot/audienceGroup/" + audienceGroupIdString
    var bearerAuth = "Bearer " + token

    // new get request
    req, err := http.NewRequest("GET", getAudiAPI, nil)
    if err != nil {
        return
    }
    req.Header.Set("Content-Type", "application/json")
    req.Header.Add("Authorization", bearerAuth)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Println("Error on response.\n[ERRO] -", err)
    }

    // get response
    var respResult map[string]interface{}
    json.NewDecoder(resp.Body).Decode(&respResult)

    if resp.StatusCode == 400 {
        err = errors.New(respResult["message"].(string))
        return
    }

    isExist = true
    audienceGroup := respResult["audienceGroup"]
    description = audienceGroup.(map[string]interface{})["description"].(string)
    return
}

// 檢查是否有此名稱(description)的audience
func (audi *CreateAudiBody) CheckAudienceByDesc(description string) (isExist bool, audienceGroupId int, err error) {
    // 取得line token
    t := Token{ID: "0"}
    token, _ := t.GetToken()

    // 取得 audience list
    listAudiAPI := "https://api.line.me/v2/bot/audienceGroup/list"
    var bearerAuth = "Bearer " + token

    // new get request
    req, err := http.NewRequest("GET", listAudiAPI, nil)
    if err != nil {
        return
    }
    req.Header.Set("Content-Type", "application/json")
    req.Header.Add("Authorization", bearerAuth)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Println("Error on response.\n[ERRO] -", err)
    }

    // get response
    var respResult map[string]interface{}
    json.NewDecoder(resp.Body).Decode(&respResult)

    if resp.StatusCode == 400 {
        err = errors.New(respResult["message"].(string))
        return
    }

    for _, item := range(respResult["audienceGroups"].([]interface{})) {
        if description == item.(map[string]interface{})["description"] {
            isExist = true
            audienceGroupId = int(item.(map[string]interface{})["audienceGroupId"].(float64))
        }
    }
    return
}

// 創立一個Aidence
func (audi *CreateAudiBody) CreateAudience(description string, audiences []AudiencesID) (audienceGroupId int, err error) {
    // 取得line token
    t := Token{ID: "0"}
    token, _ := t.GetToken()

    // 建立audience
    createAudiAPI := "https://api.line.me/v2/bot/audienceGroup/upload"
    var bearerAuth = "Bearer " + token

    // json body
    createAudiBody := CreateAudiBody{
        Description: description,
        IsIfaAudience: false,
        Audiences: audiences,
    }

    jsonBody, err := json.Marshal(createAudiBody)
    if err != nil {
        log.Println("json body error:", err)
    }

    // new post request
    req, err := http.NewRequest("POST", createAudiAPI, strings.NewReader(string(jsonBody)))
    if err != nil {
        return
    }
    req.Header.Set("Content-Type", "application/json")
    req.Header.Add("Authorization", bearerAuth)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Println("Error on response.\n[ERRO] -", err)
    }

    // get response
    var respResult map[string]interface{}
    json.NewDecoder(resp.Body).Decode(&respResult)

    if resp.StatusCode == 400 {
        err = errors.New(respResult["message"].(string))
        return
    }

    audienceGroupId = int(respResult["audienceGroupId"].(float64))
    return
}

func (audi *UpdateAudiBody) UpdateAudi(audienceGroupId int, audiences []AudiencesID) (err error) {
    // 取得line token
    t := Token{ID: "0"}
    token, _ := t.GetToken()

    // update audience
    updateAudiAPI := "https://api.line.me/v2/bot/audienceGroup/upload"
    var bearerAuth = "Bearer " + token

    // json body
    updateAudiBody := UpdateAudiBody{
        GroupID: audienceGroupId,
        Audiences: audiences,
    }

    jsonBody, err := json.Marshal(updateAudiBody)
    if err != nil {
        log.Println("json body error:", err)
    }

    // new put request
    req, err := http.NewRequest("PUT", updateAudiAPI, strings.NewReader(string(jsonBody)))
    if err != nil {
        return
    }
    req.Header.Set("Content-Type", "application/json")
    req.Header.Add("Authorization", bearerAuth)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Println("Error on response.\n[ERRO] -", err)
    }

    // get response
    var respResult map[string]interface{}
    json.NewDecoder(resp.Body).Decode(&respResult)

    if resp.StatusCode == 400 {
        err = errors.New(respResult["message"].(string))
        return
    }
    return
}

func (audi *UpdateAudiBody) DeleteAudience(audienceGroupId int) (err error) {
    // 取得line token
    t := Token{ID: "0"}
    token, _ := t.GetToken()

    // 刪除audience
    audienceGroupIdString := fmt.Sprint(audienceGroupId)
    deleteAudiAPI := "https://api.line.me/v2/bot/audienceGroup/" + audienceGroupIdString
    var bearerAuth = "Bearer " + token

    // new delete request
    req, err := http.NewRequest("DELETE", deleteAudiAPI, nil)
    if err != nil {
        return
    }
    req.Header.Set("Content-Type", "application/json")
    req.Header.Add("Authorization", bearerAuth)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Println("Error on response.\n[ERRO] -", err)
    }

    // get response
    var respResult map[string]interface{}
    json.NewDecoder(resp.Body).Decode(&respResult)

    if resp.StatusCode == 400 {
        err = errors.New(respResult["message"].(string))
        return
    }
    return
}

func (audi *Audience) GetAllAudience() (data []Audience, err error) {
    data = make([]Audience, 0)

    // 取得line token
    t := Token{ID: "0"}
    token, _ := t.GetToken()

    // 取得 audience list
    listAudiAPI := "https://api.line.me/v2/bot/audienceGroup/list"
    var bearerAuth = "Bearer " + token

    // new get request
    req, err := http.NewRequest("GET", listAudiAPI, nil)
    if err != nil {
        return
    }
    req.Header.Set("Content-Type", "application/json")
    req.Header.Add("Authorization", bearerAuth)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Println("Error on response.\n[ERRO] -", err)
    }

    // get response
    var respResult map[string]interface{}
    json.NewDecoder(resp.Body).Decode(&respResult)

    if resp.StatusCode == 400 {
        err = errors.New(respResult["message"].(string))
        return
    }

    // 解構 資料
    for _, item := range(respResult["audienceGroups"].([]interface{})) {
        audiItem := Audience{
            GroupID: int(item.(map[string]interface{})["audienceGroupId"].(float64)),
            Description: item.(map[string]interface{})["description"].(string),
            Count: int(item.(map[string]interface{})["audienceCount"].(float64)),
        }

        data = append(data, audiItem)
    }

    return
}