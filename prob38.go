package main

import (
  "fmt"
  "euler"
)

func pandigital(n int) bool {
   counts := make([]int,10)
   ds := euler.GetDigits(n)
   for _, v := range ds {
     if v == 0 {
        return false
     }
     counts[v]++
     if counts[v] > 1 {
        return false
     }
   }
   for k, v := range counts {
    if k <= len(ds) && k > 0 {
      if v == 0 { return false }
    } else {
      if v == 1 { return false }
    }
   }
   return true
}

func concat(x int, y int) int {
    return (x * 100000) + y
}

func main() {
   largest := 918273645
   for x := 9182; x < 10000; x++ {
     c := concat( x, x*2)
     if pandigital(c) {
       largest = c
     }
   }
   fmt.Println(largest)
}
