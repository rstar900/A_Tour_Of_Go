// https://go.dev/tour/moretypes/26

package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	// Variables to keep track of state of the closure
	called := false // indicates if it is called atleast once
	first := 0
	second := 1
	
	return func() int {
		if !called {
			// First call scenario
			called = true
			return first
		} else {
			// Subsequent calls
			result := first + second
			first = second
			second = result
			return result
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
