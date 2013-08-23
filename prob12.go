package main

import "fmt"

func div_no(n int) int {
   t := 0
   for i := 1; i*i < n+1; i++ {
      if ( n % i == 0 ) {
         t+=2
         if ( n / i == i ) {
           t--
         }
      }
   }
   return t
}

func main() {
   target := 500
   current := 0
   tri := 0
   for pos := 1; current < target; pos++ {
      tri+=pos
      current=div_no(tri)
      // fmt.Printf("%v : %v\n", tri, current)
   }
   fmt.Printf("%v : %v\n", tri, current)
}
      
