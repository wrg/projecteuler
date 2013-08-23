package main

import (
  "fmt"
  "strconv"
)

func longDiv(a int, b int, max int) (string,int) {
   remMap := make(map[int]int)
   resStr := ""
   hasDot := false
   p := 0
   for x := 0; x <= max; x++ {
      first := true
      for ; b > a ; a = a * 10 {
         if !hasDot {
            hasDot = true
            if len(resStr) == 0 {
               resStr += "0."
            } else {
               resStr += "."
            }
         }
         if first {
           first = false
         } else {
           resStr += "0"
           p++
         }
      }
      
      resStr += strconv.Itoa(a/b)
      a = a % b
      if remMap[a] != 0 {
        // fmt.Printf("%s Repeats from %d to %d\n", resStr, remMap[a], p)
        return resStr, p - remMap[a]
      } else {
        remMap[a] = p
      }
      if a == 0 {
         return resStr, 0
      }
      p++
   }
   return resStr, 0
}



func main() {
  numStr := ""
  cycLen := 0
  longX := 0
  for x := 2 ; x < 1000; x++ {
      s, c := longDiv(1,x,20000)
      // fmt.Println(c)
      if c > cycLen {
         cycLen = c
         numStr = s
         longX = x
      }
  }
  fmt.Println(numStr)
  fmt.Println(len(numStr))
  fmt.Println(cycLen)
  fmt.Println(longX)
}
