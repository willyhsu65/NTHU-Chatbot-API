package router

import (
    "github.com/gin-gonic/gin"
    
    "nthu-chatbot-api/router/apis/v1"
    "nthu-chatbot-api/router/apis/beta"
    "nthu-chatbot-api/pkg"
)

func InitRouter() *gin.Engine {
    r := gin.Default()

    apibeta := r.Group("/api/beta", pkg.EnableCookieSession())
    {
        admin := apibeta.Group("/admin", pkg.AuthSessionMiddle())
        {
            admin.GET("/userInfo", beta.GetUserInfoAdminApi)
        }

        msg := apibeta.Group("/msg", pkg.AuthSessionMiddle())
        {
            msg.POST("/push", beta.PushMsgApi)
        }
    }

    apiv1 := r.Group("/api/v1", pkg.EnableCookieSession())
    {
        apiv1.GET("/ping", v1.PingApi)
        apiv1.POST("/login", v1.LoginApi)
        apiv1.POST("/logout", v1.LogoutApi)
        apiv1.GET("/checkLogin", v1.CheckLoginApi)

        user := apiv1.Group("/user")
        {
            user.POST("/insertOne", v1.InsertOneUserApi)
            user.GET("/isExist", v1.IsExistUserApi)
            user.GET("/getFlag", v1.GetFlagUserApi)
            user.POST("/initFlag", v1.InitFlagUserApi)
            user.POST("/setFlag", v1.SetFlagUserApi)
            user.GET("/getInfo", v1.GetInfoUserApi)
            user.POST("/bindStudentID", v1.BindStudentIDUserApi)
            
            mapRecord := user.Group("/map")
            {
                mapRecord.POST("/insert", v1.InsertMapRecordUserApi)
                mapRecord.GET("/record", v1.GetMapRecordUserApi)
            }
        }

        bus := apiv1.Group("/bus")
        {
            bus.GET("/schedule", v1.GetScheduleBusApi)
        }

        phone := apiv1.Group("/phone")
        {
            phone.GET("/list", v1.GetPhonesApi)
        }

        andx := apiv1.Group("/andx")
        {
            andx.POST("/insertOne", v1.InsertANDXApi)
            andx.GET("/getOne", v1.GetRandANDXApi)
        }

        affair := apiv1.Group("/affair")
        {
            affair.GET("/ques", v1.FindQuesAffairApi)
        }

        feedback := apiv1.Group("/feedback")
        {
            feedback.POST("/insertOne", v1.InsertFeedbackApi)
        }

        data := apiv1.Group("/data")
        {
            data.GET("/qa", v1.GetQADataApi)
            data.GET("/epidemic", v1.GetEpidDataApi)
            data.GET("/recNews", v1.GetRecNewsDataApi)
        }

        msg := apiv1.Group("/msg")
        {
            msg.POST("/push", v1.PushMsgApi)
        }

        admin := apiv1.Group("/admin", pkg.AuthSessionMiddle())
        {
            admin.GET("/backup", v1.BackupAdminApi)
            admin.GET("/userInfo", v1.GetUserInfoAdminApi)
            admin.GET("/userNewCount", v1.GetUserNewCountAdminApi)
        }

        group := apiv1.Group("/group", pkg.AuthSessionMiddle())
        {
            group.GET("/anecdote", v1.GetAndxGroupApi)
            group.POST("/delAnecdote", v1.DeleteAndxGroupApi)
            group.GET("/feedback", v1.GetFeedbackGroupApi)
            group.POST("/delFeedback", v1.DeleteFeedbackGroupApi)
            group.POST("/addQACategory", v1.AddQACategoryGroupApi)
            group.POST("/addOneQA", v1.AddOneQAGroupApi)
            group.GET("/getAllQA", v1.GetAllQAGroupApi)
            group.POST("/delOneQA", v1.DelOneQAGroupApi)
            group.GET("/getAllEpid", v1.GetAllEpidGroupApi)
            group.POST("/updateOneEpid", v1.UpdateOneEpidGroupApi)
        }

        instanews := apiv1.Group("/instanews", pkg.AuthSessionMiddle())
        {
            instanews.POST("/post", v1.PostNewsInstanewsApi)
            instanews.GET("/getOne", v1.GetOneNewsInstanewsApi)
            instanews.GET("/news", v1.GetNewsInstanewsApi)
            instanews.POST("/updateOne", v1.UpdateOneNewsInstanewsApi)
            instanews.POST("/delete", v1.DeleteNewsInstanewsApi)
        }

        token := apiv1.Group("/token")
        {
            token.GET("/auth", v1.GetLineAuthTokenApi)
        }

        audience := apiv1.Group("/audi")
        {
            audience.POST("/create", v1.CreateAudiApi)
            audience.POST("/delete", v1.DeleteAudiApi)
            audience.POST("/bindUser", v1.BindUserAudiApi)
            audience.GET("/list", v1.GetAllAudiApi)
        }
    }

    return r
}