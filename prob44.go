package main

import (
   "fmt"
   "math"
)

func reverse_pent(n int) float64 {
   p := math.Sqrt((float64(2 * n)/float64(3)) + ( float64(1) / float64(36) )) + ( float64(1) / float64(6) )
   return p
}

func pent(n int) int {
   return (n * ((3*n)-1))/2
}

func isPent(n int) bool {
   r := int(math.Ceil(reverse_pent(n)))
   if n == pent(r) {
      return true
   } 
   return false
}

func main() {
  var plist = []int {}
  for x := 1; x < 10000; x++ {
    j := pent(x)
    plist = append(plist,j)
    if j > 12 {
    for _, k := range plist {
       if k != j {
        sum := k + j
        diff := j - k
        if isPent(sum) {
           // fmt.Printf("j = %d, k = %d, S = %d\n", j, k ,sum)
           if isPent(diff) {
           fmt.Printf("j = %d, k = %d, D = %d\n", j, k ,diff)
           }
        }
       }
    }
    }
  }
}
