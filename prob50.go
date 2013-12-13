package main

import (
   "fmt"
   "euler"
)

const MAX = 1000000

func main() {
    sieve := euler.MakeSieve(MAX)
    start := 2
    sum := 0
    l_sum := 0
    l_start := start
    l_count := 0
    count := 0
    for x := start; x < MAX; x++ {
      if !sieve[x] {
         count++
         sum = sum + x
         if sum >= MAX {
           start++
           sum=0
           if start > (MAX/2) {
              x = MAX
           } else {
              fmt.Printf("Start: %d, count: %d, sum: %d\n", l_start, l_count, l_sum)
              x=start
              count=0
           }
         }
         if !sieve[sum] && count > l_count {
            l_count = count
            l_sum = sum
            l_start = start
         }
       }
    }
}
           

