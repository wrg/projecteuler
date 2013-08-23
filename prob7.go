package main
/*
   Problem 7:
   By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

   What is the 10 001st prime number?
*/

import ( "fmt"
         "math" )

func is_prime(n int) bool {
   if ( n < 4 && n > 1 ) { return true }
   if ( n % 2 == 0 ) { return false }
   max := math.Sqrt(float64(n))
   for x := 2; float64(x) <= max ; x++ {
      if n % x == 0 {
         return false
      }
   }
   return true
}

func main() {
   pos := 0
   cur := 2
   pri :=0
   // seems inefficient to check all numbers until you find the 10001st one, but runs fast enough.
   for pos < 10001 {
      if is_prime(cur) {
         pos++
         pri = cur
      }
      cur++
   }
   fmt.Printf("%v (%v)\n",pri,pos)
}
