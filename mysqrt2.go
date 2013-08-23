package main

import (
    "fmt"
)

func Sqrt(x float64) float64 {
    z := 1.0
    last := x
    delta := last - z
    for y :=0; delta > .0001; y++ {
        last = z
        z = z - ((z*z - x)/2*z)
        delta = last - z
        if ( delta < 0 ) { delta = -delta }
        if y % 1000 == 0 { fmt.Println(z,delta) }
    }
    return z
}

func main() {
    fmt.Println(Sqrt(2))
}
