package main

import (
    "fmt"
)

func Sqrt(x float64) float64 {
    z := x / 2
    delta := z / 2
    for y := 0; delta > .000000000000001 ; y++ {
        if z * z == x {
            return z
        }
        if z * z > x {
            z-=delta
        } else {
            z+=delta
        }
        delta = delta / 2
        fmt.Printf("i: %v\n", y)
    }
    return z
}

func main() {
    fmt.Println(Sqrt(212))
}
