// https://go.dev/tour/moretypes/23

package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
    m := make(map[string]int)
	
	for _, word := range strings.Fields(s) {
	    m[word]++    
	}
	
	return m
}

func main() {
	wc.Test(WordCount)
}
