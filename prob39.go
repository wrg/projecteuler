package main

import (
   "fmt"
)


func main() {
   pmax := 0
   pscount := 0
   for p := 12; p < 1001; p++ {
      f := 0
      for a := 1; a < p; a++ {
       for b := 1; b < p - a; b++ {
          c := p - (a + b)
          if c > 0 && ((a*a) + (b*b)) == (c*c) {
            f++
          }
       }
      }
      if f > pscount {
        pscount = f
        pmax = p
      }
    }
    fmt.Println("P is:", pmax)
    fmt.Println("Number of solutions:", pscount)
}
