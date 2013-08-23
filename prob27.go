package main

import (
   "fmt"
   "math"
)

type QPrime struct {
   sieve []bool
   A int
   B int
   Prod int
   NumPrimes int
}

func (q *QPrime) make_sieve(max int) {
   q.sieve = make([]bool,max,max)
   smax := math.Sqrt(float64(max)) + 1
   for x := 2; float64(x) < smax; x++ {
      if q.sieve[x] != true {
      for y := x + x ; y < max; y+=x {
         q.sieve[y]=true
      }
      }
   }
}

func (q *QPrime) conPrimeCount(a int, b int) {
   var n int = 0
   var r int
   searching := true
   for searching {
      r = (n * n) + (a * n) + b
      if r > 0 && q.sieve[r] != true {
        n++
      } else {
        searching = false
      }
   }
   if n > q.NumPrimes {
      q.A = a
      q.B = b
      q.Prod = a * b
      q.NumPrimes = n
   }
}

func whichOp(x int) string {
  if x < 0 { return "-" }
  return "+"
}

func Abs(x int) int {
   if x < 0 { return -x }
   return x
}

func (q *QPrime) PrintFormula() string {
   s1 := whichOp(q.A)
   s2 := whichOp(q.B)
   return fmt.Sprintf("n^2 %s %dn %s %d",s1,Abs(q.A),s2,Abs(q.B))
}

func main() {
   q := new(QPrime)
   q.make_sieve(2001000)
   for a := -1000; a < 1000; a++ {
      for b := -1000; b < 1000; b++ {
          q.conPrimeCount(a,b)
      }
   }
   fmt.Printf("Max Consecutive Primes: %d (%s), Product: %d\n",q.NumPrimes,q.PrintFormula(),q.Prod)
   // fmt.Printf("Test: %d\n",q.conPrimeCount(1,41))
}
