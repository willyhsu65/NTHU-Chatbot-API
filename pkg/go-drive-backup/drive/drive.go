package drive

import (
   "log"
   "io"

   "google.golang.org/api/drive/v3"
)

func CreateDir(service *drive.Service, name string, parentId string) (*drive.File, error) {
   d := &drive.File{
      Name:     name,
      MimeType: "application/vnd.google-apps.folder",
      Parents:  []string{parentId},
   }
 
   file, err := service.Files.Create(d).Do()
   if err != nil {
      log.Println("Could not create dir: " + err.Error())
      return nil, err
   }
 
   return file, nil
}

func CreateFile(service *drive.Service, name string, content io.Reader, parentId string) (*drive.File, error) {
   f := &drive.File{
      Name:     name,
      Parents:  []string{parentId},
   }
   
   file, err := service.Files.Create(f).Media(content).Do()
   if err != nil {
      log.Println("Could not create file: " + err.Error())
      return nil, err
   }
 
   return file, nil
}