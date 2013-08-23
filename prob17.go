package main

import (
   "fmt"
   "strconv"
   "strings"
)

var ones = []string { "", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen" }
var tens = []string { "", "ten", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety" }
var hundreds = []string { "", "hundred", "thousand" }
var lc int = 0

type Numeric struct {
   Num int
   Str string
   Len int
}

func convert_array(ns []string) []int {
   nums := make([]int,len(ns))
   var err error
   for x := 0; x < len(ns); x++ {
      nums[x], err = strconv.Atoi(ns[x])
      if err != nil {
        fmt.Println("error:", err)
      }
   }
   return nums
}

func (n *Numeric) parseInt(i *int) {
  // var err error
  n.Num = *i
  n.Str = ""
  if n.Num < 20 {
     n.Str = ones[n.Num]
     n.Len = len(n.Str)
  } else {
     digits := convert_array(strings.Split(strconv.Itoa(n.Num),""))
     l := len(digits)
     if l == 4 {
        n.Str = fmt.Sprintf("%sthousand",ones[digits[0]])
     }
     if l >= 3 && digits[l-3] > 0 {
        if digits[l-2] > 0 || digits[l-1] > 0 {
          n.Str = fmt.Sprintf("%s%shundredand", n.Str, ones[digits[l-3]] )
        } else {
          n.Str = fmt.Sprintf("%s%shundred", n.Str, ones[digits[l-3]] )
        }
     }
     if l >=2 && digits[l-2] > 0 {
        if digits[l-2] == 1 {
           n.Str = fmt.Sprintf("%s%s", n.Str, ones[digits[l-1]+10] )
        } else {
           n.Str = fmt.Sprintf("%s%s%s", n.Str, tens[digits[l-2]], ones[digits[l-1]] )
        }
     } else {
        n.Str = fmt.Sprintf("%s%s", n.Str, ones[digits[l-1]] )
     }
     n.Len = len(n.Str)
  }
}

func main() {
  n := new(Numeric)
  l := 0
  for i := 1 ; i < 1001 ; i++ {
      n.parseInt(&i)
      fmt.Println(n.Str)
      l += n.Len
  }  
  fmt.Println(l)
}   
     
      
