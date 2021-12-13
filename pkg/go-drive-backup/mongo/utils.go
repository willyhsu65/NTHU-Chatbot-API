package mongo

import (
    "os"
    "log"
    "fmt"
    "path"
    "reflect"
    "encoding/json"
)

func WriteDoc2File(results interface{}, backupDirPath string, fileName string) {
    fmt.Println(results)

    switch reflect.TypeOf(results).Kind() {
    case reflect.Slice:
        reValue := reflect.ValueOf(results)
        reLen := reValue.Len()

        for i := 0; i < reLen; i++ {
            reItem := reValue.Index(i)
            log.Println(reItem)

            b, err := json.Marshal(reItem)
            if err != nil {
                log.Fatal(err)
            }
            
            f, err := os.OpenFile(path.Join(backupDirPath, fileName), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
            if err != nil {
                log.Fatal(err)
            }
            if _, err := f.Write(b); err != nil {
                log.Fatal(err)
            }
            if _, err := f.Write([]byte("\n")); err != nil {
                log.Fatal(err)
            }
            if err := f.Close(); err != nil {
                log.Fatal(err)
            }
        }
    }
}
