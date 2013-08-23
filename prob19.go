package main

import (
   "fmt"
   "time"
)

func main() {
   tot := 0
   for y := 1901 ; y < 2001 ; y++ {
      for m := 1 ; m < 13 ; m++ {
        t := time.Date(y, time.Month(m), 1, 23, 0, 0, 0, time.UTC)
        if  t.Weekday() == 0 {
           tot+=1
        }
      }
   }
   fmt.Println(tot)
}
         
