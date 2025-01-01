// https://go.dev/tour/basics/13

package main

import (
	"fmt"
	"math"
)

func main() {
	// Replacing var with short assignments
	x, y := 3, 4
	f := math.Sqrt(float64(x*x + y*y))
	z := uint(f)
	fmt.Println(x, y, z)
}
