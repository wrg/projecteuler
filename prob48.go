package main

import (
   "fmt"
   "math/big"
)

func main() {

   sum := big.NewInt(1)
   var x int64
   for x = 2; x < 1001; x++ {
      n := big.NewInt(x)
      n = n.Exp(n,n,nil)
      sum = sum.Add(sum,n)
   }
   sum_str := sum.String()
   l := len(sum_str)
   last_ten := sum_str[l-10:]
   fmt.Println(last_ten)
}
