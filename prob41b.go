package main

import (
   "fmt"
   "strconv"
   "euler"
)

type Digits map[rune]int

func pandigital(n int) bool {
   counts := make(Digits,10)
   ds := strconv.Itoa(n)
   for _, v := range ds {
     if v == '0' {
        return false
     }
     counts[v]++
     if counts[v] > 1 {
        return false
     }
   }
   maxv := strconv.Itoa(len(ds))
   for k, v := range counts {
    if k <= rune(maxv[0]) {
      if v == 0 { return false }
    } else {
      if v == 1 { return false }
    }
   }
   return true
}

func main() {
   s := euler.MakeSieve(999999999)
   largest := 1
   for k, v := range s {
      if !v {
        if pandigital(k) {
          largest = k
          fmt.Println(k)
        }
   }
   }
   fmt.Println(largest)
}
