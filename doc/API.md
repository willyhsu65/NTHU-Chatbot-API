#  啟動伺服器

- Golang Version：go1.13.4

- Use Golang Gin：`go run main.go`

- Build and start the server

  ~~~go
  go build
  ./app
  ~~~



# 資料庫 (MongoDB)

- **user**

  ~~~json
  {
    "userID": "", // line user id
    "category": "official" || "beta", // official帳號或beta帳號
    "time": "", // 加入時間
    "flag": "init", // 對話情境狀態
    "tag": [], // 屬於哪些line audience
    "studentID": "" // string, 學校學號
  }
  ~~~
  
- **user_map_record**

  ~~~json
  {
  	"userID": "", // line user id
  	"record": [{
      "location": "",
      "times": 0 // int32
    }]
  }
  ~~~
  
- **bus_schedule**

  ~~~json
  {
    "busType": "campus", // 校車類型
    "wdType": "daily", // 日類型
    "direction": "climb", // 校車方向
    "line": "red", // 校車路線
    "route": ["北校門口","綜二館","楓林小徑","人社院","台積館"], // 路線
    "schedule":["7:20","7:45","7:50","8:05","8:15","8:25","8:35","8:40","8:45","8:50","8:55","9:05","9:25","9:35","9:40","9:45","9:50","9:55","10:00","10:05","10:10","10:35","11:15","11:55"]
  }
  ~~~
  

* **anecdote**

  ~~~json
  {
    "id": 0, // int32
    "userID": "", // line userID
  	"time": "",
  	"content": ""
  }
  ~~~

* **feedback**

  ~~~json
  {
    "id": 0, // int32
    "category": "normal" || "epidemic",
  	"userID": "",
  	"time": "",
  	"content": ""
  }
  ~~~

* **instanews**

  ~~~json
  {
  	"id": 0, // int32
    "postTime": "", // 發佈日期
    "category": "", // 演講訊息, 藝術展覽, 各類活動
    "title": "",
    "date": [], // 若為一個區間: ["2020-03-18", "2020-03-20"]
    "time": "",
    "location": "",
    "sourceUrl": "",
    "imgUrl": "",
    "content": ""
  }
  ~~~


* **qa_data**

  ~~~json
  {
  	"id": 0, // int32
    "category": "epidemic", // QA的類別
    "data": [{
      "ques": "",
      "ans": ""
    }]
  }
  ~~~

* **epidemic_data**

  ~~~json
  {
  	"id": 0, // int32
    "category": "", // 防疫對象類別: borders(住宿生), isolation(居家隔離), foreign(非本國人), change2tkm(陸港澳交換), contact(聯絡管道)
    "title": "", // 防疫標題: 住宿生注意事項, 什麼樣的人必須居家隔離？, 非本國生注意事項, 陸港澳交換計畫相關, 相關資訊發布與聯絡管道
    "content": [] // 顯示在chatbot的內容
  }
  ~~~





# Restful API

- v1URL = http://nthu-chatbot.tk/api/v1

  betaURL = http://nthu-chatbot.tk/api/beta

- **Response Format**

  ```json
  resp = {
    "msg": "", // 訊息
    "result": "", // 回傳的資料
  }
  ```

- **Status Code**

  | Status Code             | Description                   |
  | ----------------------- | ----------------------------- |
  | 200 OK                  | 請求成功                      |
  | 400 Bad Request         | (客戶端錯誤) 收到無效語法     |
  | 404 Not Found           | (客戶端錯誤) 找不到請求的資源 |
  | 503 Service Unavailable | (伺服器錯誤) 服務無法使用     |

- **baseAPI**

  | API Method | API URL          | Desc              | Req Params        | Resp Result      |
  | ---------- | ---------------- | ----------------- | ----------------- | ---------------- |
  | POST       | v1URL/ping       | 測試後端是否alive |                   |                  |
  | POST       | v1URL/login      | 登入              | account, password | department, name |
  | POST       | v1URL/logout     | 登出              |                   |                  |
  | GET        | v1URL/checkLogin | 確認登入信息      |                   | department       |

* **userAPI**

  | API Method | API URL                               | Desc                         | Req Params        | Resp Result |
  | ---------- | ------------------------------------- | ---------------------------- | ----------------- | ----------- |
  | POST       | v1URL/user/insertOne                  | 新增使用者                   | userID            |             |
  | GET        | v1URL/user/isExist                    | 查看使用者是否存在           | userID            | isExist     |
  | GET        | v1URL/user/getFlag                    | 取得flag                     | userID            | flag        |
  | POST       | v1URL/user/initFlag                   | 初始化flag                   | userID            |             |
  | POST       | v1URL/user/setFlag                    | 設定flag                     | userID, flag      |             |
  | GET        | v1URL/user/getInfo                    | 取得使用者資料               | userID            | time        |
  | POST       | v1URL/user/bindStudentID              | 綁定學生學號                 | userID, studentID |             |
  | POST       | v1URL/user/map/insert                 | 新增一筆查詢校園地圖資料     | userID, location  |             |
  | GET        | v1URL/user/map/record                 | 取得此使用者查詢校園地圖紀錄 | userID            | [location]  |
  | POST       | v1URL/user/updateBroadcastTag         | 設定使用者的上次被推播的時間戳 | userID, tag     |   |
  | GET        | v1URL/user/getBroadcastAudienceIds    | 取得要被推播的使用者 id      |                   | [userID]  |
  
* **anecdoteAPI**

  | API Method | API URL              | Desc             | Req Params      | Resp Result |
  | ---------- | -------------------- | ---------------- | --------------- | ----------- |
  | POST       | v1URL/andx/insertOne | 新增一條趣事     | userID, content |             |
  | GET        | v1URL/andx/getOne    | 隨機取得一條趣事 |                 | content     |

* **busAPI**

  | API Method | API URL            | Desc         | Req Params                                                   | Resp Result          |
  | ---------- | ------------------ | ------------ | ------------------------------------------------------------ | -------------------- |
  | GET        | v1URL/bus/schedule | 動態校車時間 | time *(7:10:21)* , <br>busType *(campus)* , <br>direction *(climb/descend)* , <br>line *(red/green)* , <br>dep *(北校門口, 綜二館, 楓林小徑, 奕園停車場, 南門停車場, 人社院, 台積館)* | arriveTime, waitTime |
  
* **phoneAPI**

  | API Method | API URL          | Desc             | Req Params                              | Resp Result     |
  | ---------- | ---------------- | ---------------- | --------------------------------------- | --------------- |
  | GET        | v1URL/phone/list | 取得所有電話資料 | phone_type (空/academic/administration) | [{name, phone}] |

* **affairAPI**

  | API Method | API URL           | Desc         | Req Params | Resp Result |
  | ---------- | ----------------- | ------------ | ---------- | ----------- |
  | GET        | v1URL/affair/ques | 詢問校務資訊 | ques       | ans         |

* **feedbackAPI**

  | API Method | API URL                | Desc       | Req Params | Resp Result |
  | ---------- | ---------------------- | ---------- | ---------- | ----------- |
  | POST       | v1URL/feedback/insertOne | 使用者回饋 | userID, category(normal/epidemic), content |             |

* **dataAPI**

  | API Method | API URL             | Desc         | Req Params                                | Resp Result                                                  |
  | ---------- | ------------------- | ------------ | ----------------------------------------- | ------------------------------------------------------------ |
  | GET        | v1URL/data/qa       | 取得QA資料   | category                                  | [{ques, ans}]                                                |
  | GET        | v1URL/data/epidemic | 取得防疫資訊 | category                                  | []content                                                    |
  | GET        | v1URL/data/recNews  | 推薦學生新聞 | category (all/speech/exhibition/activity) | [{title, category, date, time, location, sourceUrl, imgUrl, content}] |

* **msgAPI**

  | API Method | API URL        | Desc     | Req Params       | Resp Result |
  | ---------- | -------------- | -------- | ---------------- | ----------- |
  | POST       | v1URL/msg/push | 推送訊息 | []to, []messages |             |

* **adminAPI**：後台共用的功能

  | API Method | API URL                  | Desc             | Req Params | Resp Result      |
  | ---------- | ------------------------ | ---------------- | ---------- | ---------------- |
  | GET        | v1URL/admin/backup       | 備份資料庫       |            |                  |
  | GET        | v1URL/admin/userInfo     | 取得使用者資料   |            | [{userID, time}] |
  | GET        | v1URL/admin/userNewCount | 取得今日新增用戶 |            | todayCount       |

* **groupAPI**：狗狗情報員們

  | API Method | API URL                   | Desc             | Req Params          | Resp Result                          |
  | ---------- | ------------------------- | ---------------- | ------------------- | ------------------------------------ |
  | GET        | v1URL/group/anecdote      | 取得所有趣聞     |                     | [{id, userID, time, content}]        |
  | POST       | v1URL/group/delAnecdote   | 刪除一則趣聞     | id                  |                                      |
  | GET        | v1URL/group/feedback      | 取得類別的回饋   | category            | [{id, userID, time, content}]        |
  | POST       | v1URL/group/delFeedback   | 刪除一則回饋     | id                  |                                      |
  | POST       | v1URL/group/addQACategory | 新增QA類別       | category            |                                      |
  | POST       | v1URL/group/addOneQA      | 新增QA資料       | category, ques, ans |                                      |
  | GET        | v1URL/group/getAllQA      | 取得所有QA       |                     | [{id, category, data:[{ques, ans}]}] |
  | POST       | v1URL/group/delOneQA      | 刪除一筆QA       | category, ques      |                                      |
  | GET        | v1URL/group/getAllEpid    | 取得所有防疫訊息 |                     | [{id, category, title, []content}]   |
  | POST       | v1URL/group/updateOneEpid | 更新一筆防疫訊息 | category, []content |                                      |
  
* **instanewsAPI**：沐報

  | API Method | API URL                   | Desc         | Req Params                                                   | Resp Result                                                  |
  | ---------- | ------------------------- | ------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
  | POST       | v1URL/instanews/post      | 新增新聞     | title, category, date, time, location, imgUrl, content       |                                                              |
  | GET        | v1URL/instanews/getOne    | 取得一筆新聞 | id                                                           | id, postTime, title, category, date, time, location, sourceUrl, imgUrl, content |
  | GET        | v1URL/instanews/news      | 取得所有新聞 | category (all/speech/exhibition/activity)                    | [{id, postTime, title, category, date, time, location, sourceUrl, imgUrl, content}] |
  | POST       | v1URL/instanews/updateOne | 更新一筆新聞 | id, postTime, title, category, date, time, location, sourceUrl, imgUrl, content |                                                              |
  | POST       | v1URL/instanews/delete    | 刪除新聞     | id                                                           |                                                              |

* **tokenAPI**

  | API Method | API URL          | Desc                      | Req Params          | Resp Result    |
  | ---------- | ---------------- | ------------------------- | ------------------- | -------------- |
  | GET        | v1URL/token/auth | 取得 LINE token & webhook | mode(official/beta) | token, webhook |
  
* **audienceAPI**

  | API Method | API URL             | Desc                 | Req Params                           | Resp Result                                     |
  | ---------- | ------------------- | -------------------- | ------------------------------------ | ----------------------------------------------- |
  | POST       | v1URL/audi/create   | 建立受眾             | description, audiences[{id}]         | audienceGroupId                                 |
  | POST       | v1URL/audi/delete   | 刪除受眾             | audienceGroupId                      |                                                 |
  | POST       | v1URL/audi/bindUser | lineID綁定受眾       | userID, description, audiences[{id}] | audienceGroupId                                 |
  | GET        | v1URL/audi/list     | 取得所有audience資料 |                                      | [{audienceGroupId, description, audienceCount}] |
  
* **betaAPI**

  | API Method | API URL                | Desc                 | Req Params       | Resp Result      |
  | ---------- | ---------------------- | -------------------- | ---------------- | ---------------- |
  | GET        | betaURL/admin/userInfo | 取得測試者使用者資料 |                  | [{UserID, Time}] |
  | POST       | betaURL/msg/push       | 推送訊息給測試者     | []to, []messages |                  |
