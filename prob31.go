// ProjectEuler.net - Problem 31 (Coin Sums)
// How many different ways can 2 pounds be made using any number of coins

package main

import (
   "fmt"
)

const PENCE = 1
const POUND = 100

var coins map[string]int

func Sum(c map[string]int) int {
  sum := 0
  for _, v := range c {
    sum += v
  }
  return sum
}

func main() {
 coins = map[string] int {
   "2P": 0,
   "1P": 0,
   "50p": 0,
   "20p": 0,
   "10p": 0,
   "5p": 0,
   "2p": 0,
   "1p": 0,
 }
 max := 200
 count := 0 
 for coins["2P"] = 0; coins["2P"] <= max; coins["2P"] += 2 * POUND {
  for coins["1P"] = 0; coins["1P"] <= max; coins["1P"] += POUND {
   for coins["50p"] = 0; coins["50p"] <= max; coins["50p"] += 50 * PENCE {
    for coins["20p"] = 0; coins["20p"] <= max; coins["20p"] += 20 * PENCE {
     for coins["10p"] = 0; coins["10p"] <= max; coins["10p"] += 10 * PENCE {
      for coins["5p"] = 0; coins["5p"] <= max; coins["5p"] += 5 * PENCE {
       for coins["2p"] = 0; coins["2p"] <= max; coins["2p"] += 2 * PENCE {
        for coins["1p"] = 0; coins["1p"] <= max; coins["1p"] += PENCE {
          s := Sum(coins)
          // fmt.Println(s)
          if s == 200 {
            fmt.Println(coins)
            count++
          }
         }
         coins["1p"] = 0
        }
         coins["1p"] = 0
        coins["2p"] = 0
        max = 201 - Sum(coins)
       }
         coins["1p"] = 0
        coins["2p"] = 0
       coins["5p"] = 0
       max = 201 - Sum(coins)
      }
         coins["1p"] = 0
        coins["2p"] = 0
       coins["5p"] = 0
      coins["10p"] = 0
      max = 201 - Sum(coins)
     }
         coins["1p"] = 0
        coins["2p"] = 0
       coins["5p"] = 0
      coins["10p"] = 0
     coins["20p"] = 0
     max = 201 - Sum(coins)
    }
         coins["1p"] = 0
        coins["2p"] = 0
       coins["5p"] = 0
      coins["10p"] = 0
     coins["20p"] = 0
    coins["50p"] = 0
    max = 201 - Sum(coins)
   }
         coins["1p"] = 0
        coins["2p"] = 0
       coins["5p"] = 0
      coins["10p"] = 0
     coins["20p"] = 0
    coins["50p"] = 0
    coins["1P"] = 0
    max = 201
  }
  fmt.Printf("Total: %d\n",count)
}
