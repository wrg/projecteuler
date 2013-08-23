package main

import ( "fmt"
         "math" )

func is_prime(n int64) bool {
   if ( n < 4 && n > 1 ) { return true }
   if ( n % 2 == 0 ) { return false }
   max := math.Sqrt(float64(n)) + 1
   var x int64
   for x = 2; float64(x) <= max ; x++ {
      if n % x == 0 {
         return false
      }
   }
   return true
}

func main() {
  
   var sum int64 = 5
   var x int64
   for x = 5 ; x < 2000000 ; x++ {
     if is_prime(x) {
      sum+=x
     }
   }
   fmt.Println(sum)
}
