func isPalindrome(x int) bool {

    if x<0{
        return false
    }
    
    var div int = 1
    for x>=10 * div{
        div*=10
    }
    
    for x!=0{
        right:=x%10
        left := x/div
            
        if right!=left{
            return false
        }
        x =(x%div)/10
        div = div/100
    }
    return true
}
