func strStr(haystack string, needle string) int {
        
    
    if needle==""{
        return 0
    }
    for i:=0;i<=len(haystack)-len(needle);i++{
       var found bool = true
        for j:=0;j<len(needle);j++{
            if haystack[i+j]!=needle[j]{
                found = false
                break
            }
            
        }
        if found{
            return i
        }
        
    }
    return -1
    
}