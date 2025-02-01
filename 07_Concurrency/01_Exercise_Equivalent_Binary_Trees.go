// https://go.dev/tour/concurrency/8

package main

import (
	"golang.org/x/tour/tree"
	"fmt" 
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	// Base case
	if t == nil {
		return
	}
	
	// Otherwise walk to from the left subtree current, and right one (inorder traversal)
	
	Walk(t.Left, ch) 	// Left Subtree
	ch <- t.Value    	// Current Subtree
	Walk(t.Right, ch)	// Right Subtree
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	// Create 2 buffered channels for 2 trees (since we know both will have exactly 10 values)
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	
	// Start the parallel execution of Walk() for both the trees in background
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	
	// Since it will be a buffered channels, reads will be in order for both channels
	for i := 0; i < 10; i++ {
		if <-ch1 != <-ch2 {
			// If any value is different, they are unequal trees
			return false
		}
	}
	
	// If we reached so far, no values were unequal
	// Hence, they are the same trees
	return true
}

func main() {
	/* For checking the Walk() implementation
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
	*/
	
	fmt.Println(Same(tree.New(1), tree.New(1))) // True
	fmt.Println(Same(tree.New(1), tree.New(2))) // False
}
