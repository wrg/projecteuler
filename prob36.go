package main

import (
  "fmt"
  "strings"
  "strconv"
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

func pbin_test(n int) bool {
  b := fmt.Sprintf("%b",n)
  bn := strings.Split(b, "")
  for i, j := 0, len(bn)-1; i < j; i, j = i+1, j-1 {
      bn[i], bn[j] = bn[j], bn[i]
  }
  rb := strings.Join(bn, "")
  if b == rb {
    return true
  }
  return false
}


func main() {
   sum := 0
   for x := 1; x < 1000000; x++ {
      if pa_test(x) && pbin_test(x) {
         sum += x
      }
   }
   fmt.Println(sum)
}

