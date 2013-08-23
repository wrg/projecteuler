package main
/*
  A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 X 99.

  Find the largest palindrome made from the product of two 3-digit numbers.
*/

import (
  "fmt"
//  "math/complex"
  "strconv"
  "strings"
)

var (
   palx int = 0
   paly int = 0
   pal int = 0
)

func pa_test(n int) bool {
  // convert n to string sn for parsing
  sn := strings.Split(strconv.Itoa(n), "")
  // Reverse sn
  for i, j := 0, len(sn)-1; i < j; i, j = i+1, j-1 {
      sn[i], sn[j] = sn[j], sn[i]
  }
  srn := strings.Join(sn, "")
  // convert reversed string srn to int rn
  rn, err := strconv.Atoi(srn)
  if err != nil {
    fmt.Println("Error converting string to int.")
  }
  // if n equals rn then we have a palindrome
  if n == rn {
     return true
  }
  return false
}

func main() {
   for x := 100; x < 1000; x++ {
      for y := 100; y < 1000; y++ {
         p := x*y
         if pa_test(p) && p > pal {
             palx = x
             paly = y
             pal = p
         }
      }
   }
   fmt.Printf("Result: %v * %v = %v\n", palx, paly, pal)
}
