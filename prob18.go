package main

import (
   "fmt"
)

var triangle = [][]int { {75},
{95,64},
{17,47,82},
{18,35,87,10},
{20,4,82,47,65},
{19,1,23,75,3,34},
{88,2,77,73,7,63,67},
{99,65,4,28,6,16,70,92},
{41,41,26,56,83,40,80,70,33},
{41,48,72,33,47,32,37,16,94,29},
{53,71,44,65,25,43,91,52,97,51,14},
{70,11,33,28,77,73,17,78,39,68,17,57},
{91,71,52,38,17,14,91,43,58,50,27,29,48},
{63,66,4,68,89,53,67,30,73,16,69,87,40,31},
{4,62,98,27,23,9,70,98,73,93,38,53,60,4,23} }

var max int = 16383
const one = '1'

type Path struct {
   Bpath string
   Total int
   Nums []int
}

func (p *Path) walk(b *string) {
   x := 0
   y := 0
   p.Bpath = *b
   p.Nums = make([]int, len(triangle))
   p.Total = triangle[x][y]
   p.Nums[0] = triangle[x][y]
   
   for i, v := range *b {
      if  v == one {
         y+=1
      }
      x+=1
      p.Total += triangle[x][y]
      p.Nums[i+1] = triangle[x][y]
   }

}

func main() {
   path := new(Path)
   biggest := new(Path)
 
   for z := 0 ; z <= max ; z++ {
      
      bpath := fmt.Sprintf( "%014b", z )
      path.walk(&bpath)
      if ( path.Total > biggest.Total ) {
         biggest.Total = path.Total
         biggest.Bpath = bpath
         biggest.Nums = path.Nums
      }
      fmt.Println(z, ":", path.Total)
   }
   fmt.Println("\nTotal: ", biggest.Total)
   fmt.Println("Path: ", biggest.Bpath)
   fmt.Println("Numbers:")
   fmt.Println(biggest.Nums)
}
