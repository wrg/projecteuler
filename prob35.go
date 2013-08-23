package main

import (
   "fmt"
   "math"
   "strconv"
)

type Primes []bool

func make_sieve(max int) Primes {
   p := make(Primes,max,max)
   smax := math.Sqrt(float64(max)) + 1
   for x := 2; float64(x) < smax; x++ {
      if p[x] != true {
      for y := x + x ; y < max; y+=x {
         p[y]=true
      }
      }
   }
   return p
}

func rotate(n int, x int) int {
   if n < 10 {
     return n
   }
   d := strconv.Itoa(n)
   rStr := fmt.Sprintf("%s%s",d[x:], d[0:x])
   r, err := strconv.Atoi(rStr)
   if err != nil {
      fmt.Println("Error: ",err)
      return n
   }
   return r
}

func allRotations(n int) []int {
   l := len(strconv.Itoa(n))
   r := make([]int,l)
   r[0] = n
   for x := 1; x < l; x++ {
       r[x] = rotate(n, x)
   }
   return r
}

func main() {
   p := make_sieve(1000000)
   found := make(map[int]bool)
   for k, v := range p {
      if !v && k > 1 && !found[k] {
        r := allRotations(k)
        pri := true
        for _, x := range r {
           if p[x] == true {
             pri = false
           }
        }
        if pri == true {
        for _, x := range r {
          found[x] = true
           fmt.Print(x, " ")
        }
        fmt.Println()
        }
      }
   }
   fmt.Println("Found: ",len(found))
   fmt.Println()
}

