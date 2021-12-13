package models

import (
    "context"
    "log"
    "time"
    "errors"
    "sort"
    
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson"

    db "nthu-chatbot-api/database"
    "nthu-chatbot-api/utils/collectionUtil"
)

type Instanews struct {
    Account string
}

type News struct {
    ID        int32    `json:"id" form:"id"`
    PostTime  string   `json:"postTime" form:"postTime"`
    Category  string   `json:"category" form:"category"`
    Title     string   `json:"title" form:"title"`
    Date      []string `json:"date" form:"date"`
    Time      string   `json:"time" form:"time"`
    Location  string   `json:"location" form:"location"`
    SourceUrl string   `json:"sourceUrl" form:"sourceUrl"`
    ImgUrl    string   `json:"imgUrl" form:"imgUrl"`
    Content   string   `json:"content" form:"content"`
}

func (i *Instanews) PostNews(n News) (err error) {
    instanewsCollect := db.MongoDatabase.Collection("instanews")

    // 查看是否有相同的title
    var result News
    filter := bson.M{"title": n.Title}
    err = instanewsCollect.FindOne(context.TODO(), filter).Decode(&result)
    
    if result.Title != "" {
        err = errors.New("新增失敗，已經有相同的新聞了")
        return
    }

    // 新增news
    id := collectionUtil.GetCollectionLastId(instanewsCollect)
    postTime := time.Now().Format("2006-01-02 15:04:05")
    
    var news News = News{
        ID: id,
        PostTime: postTime,
        Category: n.Category,
        Title: n.Title,
        Date: n.Date,
        Time: n.Time,
        Location: n.Location,
        SourceUrl: n.SourceUrl,
        ImgUrl: n.ImgUrl,
        Content: n.Content,
    }

    _, err = instanewsCollect.InsertOne(context.TODO(), news)
    if err != nil{
        log.Println(err)
    } else {
        log.Println("Successful post news")
    }
    return
}

func (i *Instanews) GetOne(id int32) (data News, err error) {
    instanewsCollect := db.MongoDatabase.Collection("instanews")

    err = instanewsCollect.FindOne(context.TODO(), bson.M{"id": id}).Decode(&data)
    if err != nil {
        log.Println(err)
    }
    return
}


func (i *Instanews) GetNews(category string) (newsSlice []News, err error) {
    newsSlice = make([]News, 0)

    categoryChMap := map[string]string{
        "speech": "演講訊息",
        "exhibition": "藝術展覽",
        "activity": "各類活動",
    }

    instanewsCollect := db.MongoDatabase.Collection("instanews")
    ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

    var cur *mongo.Cursor
    if category == "all" {
        cur, err = instanewsCollect.Find(ctx, bson.M{})
        if err != nil {
            log.Println(err) 
        }
    } else {
        cur, err = instanewsCollect.Find(ctx, bson.M{"category": categoryChMap[category]})
        if err != nil {
            log.Println(err) 
        }
    }

    for cur.Next(ctx) {
        var result News
        err := cur.Decode(&result)
        if err != nil { 
            log.Println(err) 
        }

        newsSlice = append(newsSlice, result)
    }
    if err := cur.Err(); err != nil {
        log.Println(err)
    }

    return
}

func (i *Instanews) UpdateOne(news News) (err error) {
    instanewsCollect := db.MongoDatabase.Collection("instanews")

    filter := bson.M{"id": news.ID}
    option := bson.M{
        "$set": news,
    }

    _, err = instanewsCollect.UpdateOne(context.TODO(), filter, option)
    if err != nil {
        log.Println(err)
    }
    return
}

func (i *Instanews) DeleteNews(id int32) (err error) {
    instanewsCollect := db.MongoDatabase.Collection("instanews")
    ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

    _, err = instanewsCollect.DeleteOne(ctx, bson.M{"id": id})
    if err != nil {
        log.Println(err)
    }
    return
}

// recommend to students
func (i *Instanews) GetRecNews(category string) (resultNewsSlice []News, err error) {
    newsSlice, err := i.GetNews(category)

    // filter by date
    todayDate := time.Now()
    resultNewsSlice = make([]News, 0)

    for _, news := range(newsSlice) {
        if len(news.Date)>0 && len(news.Date)==1 {
            startDate, err := time.Parse("2006-01-02", news.Date[0])
            if err != nil {
                log.Println(err.Error())
            }
            
            if startDate.After(todayDate) {
                resultNewsSlice = append(resultNewsSlice, news)
            }
        } else if len(news.Date)>0 && len(news.Date)==2 {
            // startDate, err := time.Parse("2006-01-02", news.Date[0])
            endDate, err := time.Parse("2006-01-02", news.Date[1])
            if err != nil {
                log.Println(err.Error())
            }

            if todayDate.Before(endDate) {
                resultNewsSlice = append(resultNewsSlice, news)
            }
        }
    }

    // 依照 Date 由小至大排序
    sort.Slice(resultNewsSlice,  func(i, j int) bool {
        d1, _ := time.Parse("2006-01-02", resultNewsSlice[i].Date[0])
        d2, _ := time.Parse("2006-01-02", resultNewsSlice[j].Date[0])
        return d1.Before(d2)
    })

    return
}