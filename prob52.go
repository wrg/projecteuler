package main

import (
   "fmt"
   "euler"
)

type digits map[int]int

func compare(x int, y int) bool {
   dcount := make(digits)
   d1 := euler.GetDigits(x)
   for _, v := range d1 {
      dcount[v]=1
   }
   d2 := euler.GetDigits(y)
   for _, v := range d2 {
      if dcount[v] != 1 {
        return false
      }
   }
   return true
}

func main() {
Main:
   for x := 10; ; x++ {
Range:
      for m := 2 ; m < 7 ; m++ {
         y := x * m
         if !compare(x,y) {
           break Range
         } else {
           if m == 6 {
             fmt.Println(x,x*2,x*3,x*4,x*5,x*6)
             break Main
           }
         }
      }
    }
}
