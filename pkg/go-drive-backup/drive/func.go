package drive

import (
    "os"
    "log"
    "time"
    "path"
    "io/ioutil"

    "google.golang.org/api/drive/v3"

    "nthu-chatbot-api/pkg/go-drive-backup"
)

func DriveBackupFunc() {
    t := time.Now()
    log.Println("Start drive backup mission at", t)

    service, err := GetService()
    if err != nil {
        log.Println("Drive get service error:", err)
    }

    // create folder on drive
    var folderName string = t.Format("2006-01-02")
    var parentId string = "1l6CZtb-wpBMXBMw8J-6E-CPIQOTVMbSd" // 清大Chatbot/backup
    folderId := CreateDriveBackupDir(service, folderName, parentId)
    
    // read local backup dir
    files, err := ioutil.ReadDir(goDriveBackup.DateBackupDirPath)
    if err != nil {
        log.Fatal(err)
    }

    // upload files to drive
    for _, file := range files {
        filePath := path.Join(goDriveBackup.DateBackupDirPath, file.Name())
        fileIO, err := os.Open(filePath)
        if err != nil {
            log.Panicf("Cannot open file: %v", err)
        }
        defer fileIO.Close()

        _, err = CreateFile(service, file.Name(), fileIO, folderId)
        if err != nil {
            log.Fatalf("Drive create file failed: %v", err)
        }

        log.Printf("Successfully upload %s backup file at %s", file.Name(), t.String())
    }

    log.Println("Drive backup mission complete at", t)
}

func CreateDriveBackupDir(service *drive.Service, folderName, parentId string) (folderId string){
    folder, err := CreateDir(service, folderName, parentId)
    if err != nil {
        log.Println("Drive create dir error:", err)
    }
    folderId = folder.Id
    log.Println("Successfully create drive backup directory at", time.Now())

    return
}