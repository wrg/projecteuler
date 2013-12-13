// ProjectEuler.net - Problem 42 (Sub-string divisibility)


package main

import (
   "fmt"
)

// factorial is the product of all natural numbers 1..n
func factorial(n int64) int64 {
    var x int64 = 1
    var f int64 = 1
    if n <= 1 {
      return f
    }
    for ; x <= n; x++ {
       f = f * x
    }
    return f
}

// an array of digits, with a delete function
type digits []int

func (d *digits) Del(i int) {
   r := make(digits, len(*d)-1)
   y := 0
   for x, v := range *d {
      if x != i {
         r[y] = v
         y++
      }
   }
   *d = r
}

func toInt(d []int) int64 {
   p := len(d) - 1
   val := int64(d[p])
   p--
   place := int64(10)
   for ; p >= 0; p-- {
    val = val + (int64(d[p]) * place)
    place = place * int64(10)
   }
   return val
}

   

func get_pandigital(n int64) []int {
   var di = digits {0,1,2,3,4,5,6,7,8,9}
   var r = make([]int,10)
   x := 0
   for ; len(di) > 0 ; {
      l := int64(len(di))
      chunks := factorial(l-1)
      index := int((n-1)/chunks)
      if index < 0 { index = 0 }
      r[x] = di[index]
      x++
      di.Del(index)
      n = n - (chunks * int64(index))
   }
   return r
}

func div_test(p []int) bool {
  var z = digits {0,2,3,5,7,11,13,17}
  for y := 1; y < 8; y++ {
      a := (p[y] * 100) + (p[y+1] * 10) + p[y+2]
      if a % z[y] != 0 {
         return false
      }
  }
  return true
}

func main() {
    max := factorial(10)
    sum := int64(0)
    for x := int64(1); x < max; x++ {
       p := get_pandigital(int64(x))
       if div_test(p) {
         fmt.Println(p)
         sum += toInt(p)
       }
    }
    fmt.Printf("Sum: %d\n",sum)
}
