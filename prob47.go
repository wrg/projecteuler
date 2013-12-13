package main

import (
   "fmt"
   "euler"
)

const MAX = 200000

type factors map[int]int
var primes = euler.MakePrimes(MAX)
var seive = euler.MakeSieve(MAX)

func get_factors(x int) factors {
   f := make(factors)
   num := x
Loop:
   for _, p := range primes {
      if p >= num { break Loop }
      if p > 1 {
      for ; num % p == 0 ; num = num / p {
           f[p]++
      }
      }
   }
   if num > 1 { 
   f[num]++ 
   if seive[num] {
     fmt.Println("ERROR ERROR ERROR!", num)
   }
   }
   return f
}

func isTheOne(x int) bool {
   f1 := get_factors(x)
   f2 := get_factors(x+1)
   f3 := get_factors(x+2)
   f4 := get_factors(x+3)
   // fmt.Printf("Debug: %d %d %d %d\n", len(f1), len(f2), len(f3), len(f4))
   if len(f1) > 3 && len(f2) > 3 && len(f3) > 3 && len(f4) > 3 {
      fmt.Println(f1)
      fmt.Println(f2)
      fmt.Println(f3)
      fmt.Println(f4)
      return true
   }
   return false
}

func main() {
   for x := 4; x < MAX-5 ; x++ {
      if seive[x] && seive[x+1] && seive[x+2] && seive[x+3] {
        // fmt.Println("testing ",x)
        if isTheOne(x) {
          fmt.Printf("The answer is: %d\n", x)
          break
        }
      }
   }
}
