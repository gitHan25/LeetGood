func kthDistinct(arr []string, k int) string {
    
    count := make(map[string]int)
    
    for _, char := range arr{
        count[char]++
    }
    
    for _, char := range arr{
        if count[char] == 1{
            k--
            if k == 0{
                return  char
            }
        }
    }
    return ""
}