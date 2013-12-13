package main

import (
   "fmt"
   "euler"
)

var sieve = euler.MakeSieve(10000)
var checked = make([]bool,10000,10000)

func toInt(d []int) int {
   p := len(d) - 1
   val := d[p]
   p--
   place := 10
   for ; p >= 0; p-- {
    val = val + (d[p] * place)
    place = place * 10
   }
   return val
}


func getPrimePermutations(x int) []int {
   // dcount := make(map[int]int)
   r := []int {}
   d := euler.GetDigits(x)
   /* for _, v := range d {
     dcount[v]++
     if dcount[v] > 1 { return r }
   } */
   pes := euler.Permute(d)
   for _, v := range pes {
      x := euler.Iati(v)
      if !sieve[x] && x > 1000 {
         checked[x] = true
         r = append(r,x)
      }
   }
   return r
}

func sort(a []int) []int {
   for x := 0; x < len(a); x++ {
     for y := x + 1; y < len(a); y++ {
        if a[x] > a[y] {
          a[x], a[y] = a[y], a[x]
        }
     }
   }
   return a
}

func examine(p []int) ([]int, int) {
  for k1, v1 := range p {
    for k2, v2 := range p {
      if k2 > k1 {
         distance := v2 - v1
         for k3, v3 := range p {
           if k3 > k2 {
              if (v3 - v2) == distance {
                 return []int{v1,v2,v3}, distance
              }
           }
         }
      }
     }
  }
  return nil, 0
}

func main() {
   for x := 1000; x < 10000; x++ {
      if !sieve[x] && !checked[x] {
         pp := getPrimePermutations(x)
         if len(pp) > 2 {
           found, distance := examine(sort(pp))
           if distance != 0 {
              fmt.Println("Found: ", found)
              fmt.Println("Distance: ", distance)
           }
         }
      }
   }
}
