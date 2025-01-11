// https://go.dev/tour/methods/23

package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

// Implementation of Read() for rot13Reader
func (r13 rot13Reader) Read(b []byte) (int, error) {
	for {
		bytes, err := r13.r.Read(b)
		
		// Transform the read bytes according to ROT13 rules
		// and add 13 to the relative value (b[i] - offset) % 26 between 0 and 25
		// adding the offset as 'A' or 'a'
	
		var offset byte
		
		for i := range b {
			
			if  b[i] >= 'A' && b[i] <= 'Z' {
				offset = 'A'
				
			} else if  b[i] >= 'a' && b[i] <= 'z' {
				offset = 'a'
				
			} else {
				continue
			}
			
			// In case it is an alphabet, convert
			b[i] = offset + (b[i] - offset + 13) % 26
		}
		
		return bytes, err
	}
	
	return 0, io.EOF
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
