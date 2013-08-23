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
   abun := make([]int,0,28124)
   sum_of_a := make([]bool,56247)
   for i := 12; i < 28124; i++ {
     j := div_no_sum(i)
     if j > i {
         abun = append(abun, i)
     }
   }

   for i := 0; i < (len(abun)-1); i++ {
      for j := 0; j <= i; j++ {
            ab := abun[i]+abun[j]
            sum_of_a[ab]=true
      }
   }

   sum := 0
   for i := 0; i < 28123; i++ {
      if  sum_of_a[i] != true {
         sum+=i
      }
   }

   fmt.Println(sum)
}
