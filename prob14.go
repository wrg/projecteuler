package main

import (
  "fmt"
//  "math/complex"
)

var (
   l_size int = 0
   l_start int64  = 0
   // looking bool = true
)

func get_collatz_size(n int64) int {
   size := 1
   for ; n > 1 ; {
     if n % 2 == 0 {
        n=n/2
     } else {
        n=3*n+1
     }
     size+=1
   }
   return size;
}

func print_collatz(n int64) {
   fmt.Printf("%v -> ", n)
   for ; n > 1 ; {
     if n % 2 == 0 {
        n=n/2
     } else {
        n=3*n+1
     }
     fmt.Printf("%v -> ", n)
   }
   fmt.Println("")
}

func main() {
   s := 0
   var x int64 = 0
   for ; x < 999999; x++ {
       s = get_collatz_size(x)
       if ( s > l_size ) {
         l_size = s
         l_start = x
       }
   }
   fmt.Printf("Result: %v (%v)\n", l_start, l_size)
   print_collatz(l_start)
}
