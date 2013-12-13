
package main

import (
   "fmt"
)

func tri(n int) int {
   return (n * (n + 1))/2
}

func pent(n int) int {
   return (n * ((3*n)-1))/2
}

func hex(n int) int {
   return (n * ((2*n)-1))
}

func main() {
Loop:
   for n := 286;; n++ {
     t := tri(n)
     for o := 165; o < n; o++ {
       p := pent(o)
       if t == p {
         for x := 143; x < o; x++ {
           h := hex(x)
           if p == h {
             fmt.Println(t)
             break Loop
           }
         }
       }
     }
   }
}
