package main

import (
   "fmt"
)

type Fraction struct {
   N int
   D int
}

func (f *Fraction) Mul(t Fraction) {
   f.N = f.N * t.N
   f.D = f.D * t.D
}

func (f *Fraction) Flip() {
   f.N, f.D = f.D, f.N
}

func (f *Fraction) Big() bool {
  if ( f.N > f.D ) {
     return true
  }
  return false
}

func (f *Fraction) Reduce() {
   if f.N > 1 {
   for x := 2; x <= f.N; x++ {
      if f.N % x == 0 && f.D % x == 0 {
         f.N = f.N / x
         f.D = f.D / x
         x = 1
      }
   }
   }
}

func (f *Fraction) String() string {
   return fmt.Sprintf("%d / %d", f.N, f.D )
}

func (f *Fraction) Equalto(t Fraction) bool {
    if t.Big() { return false }
    t2 := Fraction{f.N,f.D}
    t2.Reduce()
    t.Reduce()
    if t.N == t2.N && t.D == t2.D {
      return true
    }
    return false
}

func Permutate(n int, d int) []Fraction {
   p := 0
   t := Fraction{n,d}
   t2 := Fraction{0,0}
   r := make([]Fraction,10)
   for x := 1; x < 10; x ++ {
       t2.N = x * 10 + n
       t2.D = d * 10 + x
       // fmt.Println(t2.String())
       if t.Equalto(t2) {
         r[p] = t2
         p++
       }
       t2.N = n * 10 + x
       // fmt.Println(t2.String())
       if t.Equalto(t2) {
         r[p] = t2
         p++
       }
       t2.D = x * 10 + d
       // fmt.Println(t2.String())
       if t.Equalto(t2) {
         r[p] = t2
         p++
       }
       t2.N = x * 10 + n
       // fmt.Println(t2.String())
       if t.Equalto(t2) {
         r[p] = t2
         p++
       }
   }
   return r
}

func main() {
   f := Fraction{19,95}
   f2 := Fraction{1,5}
   fmt.Println(f.String())
   fmt.Println(f2.String())
   if f.Equalto(f2) {
     fmt.Printf("%s and %s are equal.\n",f.String(),f2.String())
   } else {
     fmt.Printf("%s and %s are NOT equal.\n",f.String(),f2.String())
   }
   s := Fraction{1,1}
   for d := 2; d < 10; d++ {
      for n := 1; n < 10; n++ {
        if n != d && n < d {
          f = Fraction{n,d}
          fmt.Printf("F: %s\n",f.String())
          pa := Permutate(n,d)
          for _, p := range pa {
          if p.N > 0 {
             fmt.Printf("Found: %s\n",p.String())
             // p.Reduce()
             s.Mul(p)
          }
          }
        }
       }
    }
    fmt.Println(s.String())   
    s.Reduce()
    fmt.Println(s.String())

}
  
