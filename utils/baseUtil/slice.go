package baseUtil

// 檢查 array 是否包含 ele
func Contains(arr []string, str string) bool {
    for _, a := range arr {
         if a == str {
                return true
         }
    }
    return false
}

// 刪除 array 指定 ele
func DeleteEle(arr []string, str string) (result []string) {
    for _, a := range arr {
        if a != str {
            result = append(result, a)
        }
    }
    return
}