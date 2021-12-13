package main

import (
    "github.com/robfig/cron"

    "nthu-chatbot-api/router"

    "nthu-chatbot-api/pkg/go-drive-backup"
    "nthu-chatbot-api/pkg/go-drive-backup/drive"
    "nthu-chatbot-api/pkg/go-drive-backup/mongo"
)

func main() {
     // set cron mission
     c := cron.New() 
     c.AddFunc("1 59 23 * * ?", goDriveBackup.CreateLocalBackupDir)
     c.AddFunc("3 59 23 * * ?", mongo.LocalBackupFunc)
     c.AddFunc("10 59 23 * * ?", drive.DriveBackupFunc)
     c.Start()

    router := router.InitRouter()
    router.Run()   
}