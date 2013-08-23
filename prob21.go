package main

import (
   "fmt"
   "math"
)

func div_no_sum(n int) int {
   s := 1
   max :=  math.Sqrt(float64(n))
   if n > 3 {
     for x :=2; float64(x) <= max ; x ++ {
       if  n % x == 0 {
          s += x
          r := n / x
          if x != r {
            s += r
          }
       }
     }
   }
   return s
}

func main() {
   ami := make([]bool,10000)
   sum := 0
   for i := 2; i < 10000; i++ {
      if  ami[i] != true {
         a := div_no_sum(i)
         if a != i {
           b := div_no_sum(a)
           if b == i {
             sum += (a + b)
             ami[a] = true
             ami[b] = true
             fmt.Printf("%d : %d (%d)\n",a,b,sum)
           }
         }
      }
   }

   fmt.Println(sum)
}
