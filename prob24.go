// ProjectEuler.net - problem 24 (Lexographic permutations)
// What is the millionth lexicographic permutation of the digits 0, 1, 2, 3, 4, 5, 6, 7, 8 and 9?

package main

import (
   "fmt"
)


// factorial is the product of all natural numbers 1..n
func factorial(n int64) int64 {
    var x int64 = 1
    var f int64 = 1
    if n == 0 {
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

func main() {
   var di = digits {0,1,2,3,4,5,6,7,8,9}
   var find int64 = 1000000
   for  ; len(di) > 0 ; {
      l := int64(len(di))
      chunks := factorial(l-1)
      index :=  int((find-1)/chunks)
      fmt.Print(di[index])
      find = find - (chunks * int64(index))
      di.Del(index)
   }
   fmt.Println("\n")
}
