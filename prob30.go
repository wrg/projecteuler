// ProjectEuler.net - Problem 30 (Digit fifth Powers)
// Find the sum of all the numbers that can be written as the sum of fifth powers of their digits.
// I think max will be (9^5) * 5 or 295245

package main

import (
   "fmt"
   "strconv"
   "strings"
)

func Pow(x int, y int) int {
   r := x
   for i := 1; i < y; i++ {
      r = r * x
   }
   return r
}

func SSPow(i int) bool {
   str := strings.Split(strconv.Itoa(i),"")
   sum := 0
   for _, ds := range str {
     d, err := strconv.Atoi(ds)
     if err != nil {
        fmt.Printf("Something strange happened: %g\n",err)
     }
     sum += Pow(d,5)
   }
   if sum == i {
     return true
   }
   return false
}
 

func main() {
   n := 0
   for a := 2; a <= 295245; a++ {
      if SSPow(a) {
         fmt.Println(a)
         n += a
      }
   }

   fmt.Printf("Sum is: %d\n",n)
}
