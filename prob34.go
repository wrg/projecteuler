// ProjectEuler.net - problem 34

package main

import (
  "fmt"
  "strconv"
  "strings"
)

// factorial is the product of all natural numbers 1..n
func factorial(n int) int {
    var x int = 1
    var f int = 1
    if n == 0 {
      return f
    }
    for ; x <= n; x++ {
       f = f * x
    }
    return f
}

func convert_array(ns []string) []int {
   nums := make([]int,len(ns))
   var err error
   for x := 0; x < len(ns); x++ {
      nums[x], err = strconv.Atoi(ns[x])
      if err != nil {
        fmt.Println("error:", err)
      }
   }
   return nums
}

func get_digits(n int) []int {
   digits := convert_array(strings.Split(strconv.Itoa(n),""))
   return digits
}

func sum_fact(n int, facts []int) int {
   d := get_digits(n)
   sum := 0
   for _, v := range d {
     sum += facts[v]
   }
   return sum
}

func main() {
   facts := make([]int,10)
   for x := 0; x < 10; x++ {
     facts[x] = factorial(x)
   }
   sum := 0
   for z := 3; z < factorial(10); z++ {
      sf := sum_fact(z, facts)
      if sf == z {
         sum += z
         fmt.Printf("%d found\n",sf)
      }
   }
   fmt.Printf("sum: %d\n",sum)
}
