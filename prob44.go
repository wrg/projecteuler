package main

import (
   "fmt"
)

var Pentagonal []bool

func MakePent(s int) {
   s = s * 2
   Pentagonal = make([]bool,(s * (3 * s + 1))/2)
   for x := 1;x < s; x++ {
      p := (x * (3 * x + 1))/2
      Pentagonal[p]=true
   }
}

func abs(n int) int {
  if n < 0 {
   return n - (n * 2)
  }
  return n
}

func main() {
   max := 10000
   MakePent(max)
   for j := 1; j < max; j++ {
     pj := (j * (3 * j + 1))/2
     for k := j; k < max; k++ {
      pk := (k * (3 * k + 1))/2
      if Pentagonal[pj + pk] && Pentagonal[abs(pk - pj)] {
         fmt.Printf("%d - %d = %d\n",pk,pj,pk-pj)
         // k, j = max, max 
      }
     }
   }
}
