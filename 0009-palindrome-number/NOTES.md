# Palindrome number (use math) 

``` golang

func isPalindrome(x int) bool {

    if x<0{
        return false
    }
    
    var div int = 1
  
    for x>=10 * div{
        div*=10
    }
    // ex: the input is 1211
   // so the div value is 1000

   for x!=0{
        right:=x%10
        left := x/div
            
        if right!=left{
            return false
        }
   // x untuk memisahkan digit kanan terakhir
        x =(x%div)/10

    // digunakan untuk mendapatkan digit paling kiri
    // ex 121, 121 / 100 = 1
        div = div/100
    }
    return true
}
```


