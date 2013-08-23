package main
/*
   The prime factors of 13195 are 5, 7, 13 and 29.

   What is the largest prime factor of the number 600851475143 ?
*/

import (
    "fmt"
)

var (
   num int64
   x int64
)

func main() {
    num = 600851475143
    x = 2
    for x < num {
      if num % x == 0 {
        fmt.Println(num,"/",x)
        num=num/x
        x=2
      } else {
        x=x+1
      }
   }
   fmt.Println("Largest prime is ",num)
}
