package main

import (
   "fmt"
   "euler"
    "strconv"
)

var p euler.Primes = euler.MakeSieve(1000000)

func rtrim(s string) string {
   if len(s) < 2 {
     return ""
   }
   return s[:len(s)-1]
}

func ltrim(s string) string {
   if len(s) < 2 {
     return ""
   }
   return s[1:len(s)]
}

func truncatable(n int) bool {
   t := true
   sn := strconv.Itoa(n)
   for sn != "" {
      x, err := strconv.Atoi(sn)
      if err != nil { fmt.Println(err) }
      if p[x] { t = false }
      //fmt.Println(sn)
      sn = rtrim(sn)
   }
   sn = strconv.Itoa(n)
   for sn != "" {
      x, err := strconv.Atoi(sn)
      if err != nil { fmt.Println(err) }
      if p[x] { t = false }
      // fmt.Println(sn)
      sn = ltrim(sn)
   }
   return t
}
   

func main() {
   count := 0
   sum := 0
   for k, v := range p {
      if !v && k > 11 {
        if truncatable(k) {
           sum += k
           fmt.Println("found ",k)
           count++
         }
      }
   }
   fmt.Println("count: ",count)
   fmt.Println("sum: ", sum)
}
