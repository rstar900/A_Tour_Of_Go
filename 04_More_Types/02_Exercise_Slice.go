// https://go.dev/tour/moretypes/18

package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
    // Allocate the outer array (rows)
    img := make([][]uint8, dy)
	
    for y := 0; y < dy; y++ {
	  // Allocate inner array (columns) for each row
		img[y] = make([]uint8, dx)
		  // Uncomment one of these functions below and
      // comment rest to see the magic (3rd one is             
      // default in this case)
	    for x := 0; x < dx; x++ {
		    //img[y][x] = (uint8(x)+uint8(y))/2
			//img[y][x] = uint8(x)*uint8(y)
			img[y][x] = uint8(x)^uint8(y)
		}
	}
	
	return img
}

func main() {
	pic.Show(Pic)
}
