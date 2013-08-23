package main

import (
  "fmt"
  "euler"
)

func main() {
   f := make([]int,1000002)
   for x := 1 ; x < 10; x++ {
      f[x] = x
      // fmt.Print(x)
   }
   p := 10
   for x := 10; p < 1000001; x++ {
       di := euler.GetDigits(x)
       for y := 0; y < len(di); y++ {
         if p < 1000001 {
           f[p] = di[y]
         }
         p++
       }
   }
   tot := 1
   fmt.Println()
   for x := 1; x < 1000001; x = x * 10 {
      tot = tot * f[x]
      fmt.Println(f[x])
   }
   fmt.Println(tot) 
}
