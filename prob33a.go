package main

import (
   "fmt"
   "euler"
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
   for x := 2; x < 10; x++ {
      if f.N % x == 0 && f.D % x == 0 {
         f.N = f.N / x
         f.D = f.D / x
         x = 1
      }
      if f.N == 1 || f.D == 1 {
         x = 10
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

func (f *Fraction) Strip() {
  nd := euler.GetDigits(f.N)
  dd := euler.GetDigits(f.D)
  switch {
  case nd[0] == dd[0]:
    f.N, f.D = nd[1], dd[1]
  case nd[0] == dd[1]:
    f.N, f.D = nd[1], dd[0]
  case nd[1] == dd[0]:
    f.N, f.D = nd[0], dd[1]
  case nd[1] == dd[1]:
    f.N, f.D = nd[0], dd[0]
  }
}

func has_potential(n int, d int) bool {
  if n >= d || d % 10 == 0 || n % 10 == 0 {
    return false
  }
  nd := euler.GetDigits(n)
  dd := euler.GetDigits(d)
  for _, nv := range nd {
   for _, dv := range dd {
     if nv == dv {
      return true
     }
  }
  }
  return false
}


func main() {
   s := Fraction{1,1}
   for d := 10; d < 100; d++ {
      for n := 10; n < 100; n++ {
        if has_potential(n,d) {
          f := Fraction{n,d}
          fmt.Println("F:",f.String())
          f2 := f
          f2.Strip()
          if f.Equalto(f2) {
            s.Mul(f)
            fmt.Println("F2:",f2.String())
          }    
        }
       }
    }
    fmt.Println(s.String())   
    s.Reduce()
    fmt.Println(s.String())
}
  
