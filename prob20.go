package main

import (
   "fmt"
   "math/big"
   "strconv"
   "strings"
)

// factorial is the product of all natural numbers 1..n
func factorial(n int64) *big.Int {
    var x int64 = 1
    f := big.NewInt(x)
    if n == 0 {
      return f
    }
    xb := new(big.Int)
    for ; x <= n; x++ {
       xb.SetInt64(x)
       f.Mul(f,xb)
    }
    return f
}

func main() {
    sum := 0
    x := factorial(100)
    xstr := strings.Split(x.String(),"")
    for _, v := range xstr {
      i, err := strconv.Atoi(v)
      if err != nil {
        fmt.Println("error: ", err)
      } else {
        sum+=i
      }
   }
   fmt.Println("Sum: ", sum)
}
