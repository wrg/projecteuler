package main

import (
   "fmt"
   "math/big"
   "strings"
   "strconv"
)

func main() {
   p := big.NewInt(int64(1000))
   bignum := big.NewInt(int64(2))
   bignum.Exp(bignum,p,nil)
   nstr := strings.Split(bignum.String(),"")
   res := 0
   for _, v := range nstr {
      x, err := strconv.Atoi(v)
      if err != nil {
       fmt.Println("Error converting string to int.")
      }
      res+=x
   }
   fmt.Println(res)
}

      
   
  
