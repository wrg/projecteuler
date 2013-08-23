// ProjectEuler.net - problem 28 (Number spiral diagonals)
// What is the sum of the numbers on the diagonals in a 1001 by 1001 spiral.

package main

import (
   "fmt"
)

// grid data type
type grid [][]int

// order for spiraling right
const RIGHT = 0
const DOWN = 1
const LEFT = 2
const UP = 3
// const sporder = []uint8{'r','d','l','u'}

type Spiral struct {
   Grid grid
   X, Y, Size, TurnDist int
   Direction int
}

func (s *Spiral) MakeGrid(size int) {
   s.Grid = make([][]int, size)
   for i := range s.Grid {
      s.Grid[i] = make([]int, size)
   }
   s.Size = size
   s.X = ((size + 1) / 2) - 1
   s.Y = s.X
   s.TurnDist = 1
   s.Direction = 0
}

func (s *Spiral) Fill() {
   max := s.Size * s.Size
   nextTurn := 1
   for x := 1 ; x <= max; x++ {
      s.Grid[s.X][s.Y] = x
      //fmt.Printf("Debug: %d [%d][%d]\n",x,s.X,s.Y)
      switch s.Direction {
         case RIGHT:
           s.Y += 1
         case LEFT:
           s.Y -= 1
         case DOWN:
           s.X += 1
         case UP:
           s.X -= 1
      }
      if nextTurn == s.TurnDist {
         nextTurn = 1
         s.Direction++
         if s.Direction > 3 {
           s.Direction = 0
         }
         if s.Direction == LEFT || s.Direction == RIGHT {
           s.TurnDist += 1
         }
      } else {
         nextTurn++
      }
   }
}

func (s *Spiral) Print() {
   for x := 0; x < s.Size; x++ {
     for y := 0; y < s.Size; y++ {
        fmt.Printf("%3d ",s.Grid[x][y])
     }
     fmt.Println()
   }
}

func (s *Spiral) DiagSum() int {
   sum := 0
   max := s.Size - 1
   for x := 0; x <= max; x++ {
     sum += s.Grid[x][x] + s.Grid[x][max-x]
   }
   sum--
   return sum
}

func main() {
   spiral := Spiral{}
   spiral.MakeGrid(19)
   spiral.Fill()
   if spiral.Size < 20 {
     spiral.Print()
   }
   fmt.Println(spiral.DiagSum())
}
