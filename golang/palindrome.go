func isPalindrome(x int) bool {
    var reverse int = 0
    var y int = x
    var rmd int
    var result bool
    for  {
       
        rmd = x%10
        reverse = reverse*10 + rmd
        x /= 10
        if (x==0) {
            break
        }
         
    }
    
    if ( y < 0){
        result = false
    } else if (y == reverse){
         result = true
    }else {
         result = false
    }   
        
     return result
    }
