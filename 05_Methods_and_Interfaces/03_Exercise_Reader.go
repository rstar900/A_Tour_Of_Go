// https://go.dev/tour/methods/22

package main

import (
	"golang.org/x/tour/reader"
	"fmt"
)

type MyReader struct{}

// Error type to denote empty slice
type EmptyBufferErr struct{}

// Implement Error interface for EmptyBufferErr
func (e EmptyBufferErr) Error() string {
	return fmt.Sprint("No space in buffer to read!")
}

// Add a Read([]byte) (int, error) method to MyReader.
func (mr MyReader) Read(b []byte) (int, error) {
	// Only give error when given slice is nil
	if b == nil {
		return 0, EmptyBufferErr{}
	} 
	
	bytes := 0 // Bytes read
	
	// Read 'A' into all the available bytes of the slice
	for i := range b {
		b[i] = 'A'
		bytes++
	}
	
	return bytes, nil
}

func main() {
	reader.Validate(MyReader{})
}
