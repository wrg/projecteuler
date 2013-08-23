package main

import (
   "aotx/csvsplit"
   "fmt"
   "os"
   "io"
   "bufio"
   "sort"
   "math/big"
)

func get_weight(p int, n string) *big.Int {
   r := 0
   for _, s := range n {
      r += ( int(s) - 64 )
   }
   r = r * p
   return big.NewInt(int64(r))
}

func main() {
   total := big.NewInt(0)
   fi, err := os.Open("names.txt")
   if err != nil { panic(err) }
   defer fi.Close()
   r := bufio.NewReader(fi)
   line, err := r.ReadString('\n')
   if err != nil && err != io.EOF { panic(err) }
   if line != "" {
      if line[len(line)-1] == '\n' {
         line = line[:len(line)-1]
      }
      names, err := csvsplit.Split(line)
      if err != nil {
         fmt.Println(err)
      }
      sort.Strings(names)
      for i, k := range names {
        total.Add(total,get_weight(i+1, k))
      }
   }
   fmt.Println(total.String())
}

   
