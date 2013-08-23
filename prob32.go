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

func main() {
   sum := 0
   found := make(map[int]bool)
   for x := 10; x < 99; x++ {
     for y := 100; y < 988; y++ {
       if x * y < 10000 {
          s := (x * 10000000) + (y * 10000) + (x*y)
          if pandigital(s) {
            p := x*y
            if !found[p] {
              found[p]=true
              fmt.Println("a:",p)
              sum += p
            } else {
              fmt.Println("dup:",p)
            }
          }
       }
      }
   }
   for x := 1000; x < 9877; x++ {
     for y := 1; y < 10; y++ {
       if x * y < 10000 {
          s := (y * 100000000) + (x * 10000) + (x*y)
          if pandigital(s) {
            p := x*y
            if !found[p] {
              found[p]=true
              fmt.Println("b:",p)
              sum += p
            } else {
              fmt.Println("dup:",p)
            }
          }
       }
      }
   }

   fmt.Println("sum: ",sum)
}
