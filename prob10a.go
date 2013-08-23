package main

import ( "fmt"
         "math"
)

func get_prime_sum(max int64) int64 {
   sieve := make([]bool,max,max)
   smax := math.Sqrt(float64(max)) + 1
   var x, y int64
   for x = 2; float64(x) < smax; x++ {
      if sieve[x] != true {
      for y = x + x ; y < max; y+=x {
         // fmt.Printf("marking %v\n",y)
         sieve[y]=true
      }
      }
   }
   var sum int64 = 0
   for x = 2; x < max; x++ {
      if sieve[x] != true {
         // fmt.Printf("Prime: %v\n",x)
         sum+=x
      }
   }
   return sum
}

func main() {
  p := get_prime_sum(2000000)
  fmt.Println(p)
}
     
