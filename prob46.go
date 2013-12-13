package main

import (
   "fmt"
   "euler"
   "math"
)

func test_gb(x int, p int) bool {
    tsq := int(math.Ceil(math.Sqrt(float64((x - p)/2))))
    try := p + (2 * (tsq * tsq))
    if x == try {
      return true
    }
    return false
}

func main() {
   s := euler.MakeSieve(9000)
Search:
   for x := 3; x < 9000; x+=2 {
   if s[x] {
      found := false
DoMath:
      for p, i := range s {
         if p > x { break DoMath }
         if !i {
            if test_gb(x, p) {
              found = true
              break DoMath
            }
         }
      }
      if !found {
        fmt.Printf("%d FAILED.\n", x)
        break Search
      } else {
        fmt.Printf("%d passed.\n", x)
      } 
    }
    }
}
