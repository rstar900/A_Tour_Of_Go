// https://go.dev/tour/generics/2

package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
	// Added extra variable to keep count of contents
	// and checking for empty list
	// We always read the size of the head element
	size uint
}

// Implementation of functions for a list

// Print contents stored inside a list by implementing
// String() function which allows it to be printed
// by functions like fmt.Printf(), fmt.Println(), etc.
func (l List[T]) String() string {
	
	// Print that the list is empty if it does not contain any elements yet
	if l.size == 0 {
		return fmt.Sprintf("<Empty List>")
	}

	// Otherwise iterate through the end, printing all contained values
	// until the next is a nil value
	output := "[ "

	for ptr := &l; ptr != nil; ptr = ptr.next {
		output += fmt.Sprint(ptr.val, " ")
	}

	return output + "]"
}

// Add an element at the end of the list
func (l *List[T]) push(val T) {
	
	// For an empty list, just put the value in val
	// and make size as 1
	if l.size == 0 {
		l.val = val
		l.size = 1
		// Otherwise, traverse till the end and make and allocate List element there
	} else {
		var ptr *List[T] 
		for ptr = l; ptr.next != nil; ptr = ptr.next {}
		e := List[T]{nil, val, 0}
		ptr.next = &e
		l.size++
	}
}

// Remove an element from the end of the list
func (l *List[T]) pop() {
	
	// If the list is empty, no need to remove anything
	if l.size == 0 {return}
	
	// Otherwise, traverse till the end and make the next pointer nil
	var ptr1 = l
	var ptr2 = l
	// ptr1 will always point to last element
	// followed by ptr2 which points the element before the last one
	for ; ptr1.next != nil; ptr1 = ptr1.next {
		ptr2 = ptr1
	}
	ptr2.next = nil
	l.size--
}

func main() {

	// Declaring 2 empty lists containing different types
	var l1 List[string]
	var l2 List[int]

	fmt.Println("Contents of l1:", l1)
	fmt.Println("Contents of l2:", l2)

	fmt.Println()

	// Add 2 values to each of the lists
	l1.push("Hola")
	l1.push("Amigos")
	l2.push(1)
	l2.push(2)

	fmt.Println("Contents of l1:", l1)
	fmt.Println("Contents of l2:", l2)
	
	fmt.Println()
	
	// Keep removing elements from both lists
	l1.pop()
	l2.pop()
	fmt.Println("Contents of l1:", l1)
	fmt.Println("Contents of l2:", l2)
	
	fmt.Println()
	
	l1.pop()
	l2.pop()
	fmt.Println("Contents of l1:", l1)
	fmt.Println("Contents of l2:", l2)
	
	fmt.Println()
	
	// Try removing from the empty lists
	l1.pop()
	l2.pop()
	fmt.Println("Contents of l1:", l1)
	fmt.Println("Contents of l2:", l2)
}
