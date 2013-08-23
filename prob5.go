package main
/*
  2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

  What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/

import (
  "fmt"
//  "math/complex"
)

var (
   // start with 20
   x int = 20
   big int = 0
)

func div_test(n int) bool {
   result := true
   for i := 2; i < 21; i++ {
      if n % i != 0 {
         result = false
         break
      }
   }
   return result
}

func main() {
   for {
      if div_test(x) {
          big=x
          break
      }
      // has to be divisible by 20 so only check multiples of 20.
      x = x + 20
   }
   fmt.Printf("Result: %v\n", big)
}
