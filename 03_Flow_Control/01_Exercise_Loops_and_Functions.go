// https://go.dev/tour/flowcontrol/8

package main

import (
	"fmt"
	//"math" // Uncomment to compare with function from math package
)

// Helper function to calculate absolute values
func abs(x float64) float64 {
	if x < 0 {
		return -1 * x
	}
	return x
}

func Sqrt(x float64) float64 {
	// This variable will be used to return the number whose square is closest to x
	z := 1.0
	// Old value of z to use for stopping condition (initialize to a different value than z
	// so that the loop runs atleast once)
	old_z := 0.0
	
	// Stop the adsjustments when delta is less than or equal to 0.000000000000001 (1e-15)
	for i := 0; abs(old_z - z) > 1e-15; i++ {
		//fmt.Println("Iteration", i, ":", z);
		// Store current value of z in old_z for comparison of delta
		old_z = z;
		// The adjustment to bring z closer to desired result
		z -= (z*z - x) / (2*z)
	}
	
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	//fmt.Println(math.Sqrt(2)) // Uncomment to compare with function from math package
}
