package goDriveBackup

import (
    "os"
    "log"
    "time"
    "path"
)

var LocalBackupDirPath string
var DateBackupDirPath string

func CreateLocalBackupDir() {
    pwd, _ := os.Getwd()
    t := time.Now()
    date := t.Format("2006-01-02")

    LocalBackupDirPath = path.Join(pwd, "backup")
    if _, err := os.Stat(LocalBackupDirPath); os.IsNotExist(err) {
        os.Mkdir(LocalBackupDirPath, os.ModePerm)
        log.Println("Backup directory not exist. Sucessfully create backup directory at", t)
    }

    DateBackupDirPath = path.Join(LocalBackupDirPath, date)
    if _, err := os.Stat(DateBackupDirPath); os.IsNotExist(err) {
        os.Mkdir(DateBackupDirPath, os.ModePerm)
        log.Println("Successfully create local backup directory at", t)
    }
}