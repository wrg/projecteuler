package main

import (
   "aotx/csvsplit"
   "fmt"
   "os"
   "io"
   "bufio"
)

var TriNum []bool = make([]bool,50000)

func MakeTri() {
   for x := 1; x < 265; x++ {
      t := ((x * x) + x) / 2
      TriNum[t] = true
   }
}

func StringVal(n string) int {
   r := 0
   for _, s := range n {
      r += ( int(s) - 64 )
   }
   return r
}

func main() {
   count := 0
   MakeTri()
   fi, err := os.Open("words.txt")
   if err != nil { panic(err) }
   defer fi.Close()
   r := bufio.NewReader(fi)
   line, err := r.ReadString('\n')
   if err != nil && err != io.EOF { panic(err) }
   if line != "" {
      if line[len(line)-1] == '\n' {
         line = line[:len(line)-1]
      }
      words, err := csvsplit.Split(line)
      if err != nil {
         fmt.Println(err)
      }
      for _, k := range words {
        wordVal := StringVal(k)
        if TriNum[wordVal] {
          count++
          fmt.Println(k,wordVal)
        }
      }
   }
   fmt.Println(count)
}

   
