package main
/*
   Problem 15:
     Starting in the top left corner of a 2X2 grid, and only being able to move to the right and down, there are exactly 6 routes
     to the bottom right corner.

     How many such routes are there through a 20X20 grid?
*/

import (
  "fmt"
  "math/big"
)

var (
   x_size int64 = 20
   y_size int64 = 20
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

// binomial coefficient is the number of combinations of k items that can be selected from a set of n items.
// also useful in finding the number of paths.
func binomial_coefficient(n int64, k int64) *big.Int {
   // formula is n! / ( (n-k)! * k!
   num := new(big.Int)
   num.Add(num,factorial(n))
   den := new(big.Int)
   den.Add(den,factorial(n-k))
   den.Mul(den,factorial(k))
   return num.Div(num, den)
}

func main() {
   // fmt.Println(factorial(x_size+y_size))
   fmt.Println(binomial_coefficient(x_size+y_size,x_size))
}
