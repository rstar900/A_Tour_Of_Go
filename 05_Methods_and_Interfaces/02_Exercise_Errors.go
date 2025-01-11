// https://go.dev/tour/methods/20

package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

// Implement Error() function for ErrNegativeSqrt to make it compatible wih Error interface
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprint("cannot Sqrt negative number: ", float64(e))
}

// Helper function to calculate absolute values
func abs(x float64) float64 {
	if x < 0 {
		return -1 * x
	}
	return x
}

func Sqrt(x float64) (float64, error) {
	
	if x < 0 {
		// In this case we return the error with a non nil value and float64 value as 0
		return 0, ErrNegativeSqrt(-2)
	}
	
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
	
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
