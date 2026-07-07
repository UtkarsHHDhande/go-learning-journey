// /Exercise: Maps
// Implement WordCount. It should return a map of the counts of each “word” in the string s. The wc.Test function runs a test suite against the provided function and prints success or failure.

package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	count := make(map[string]int)

	for _, word := range words {
		count[word]++
	}

	return count
}

func main() {
	wc.Test(WordCount)
}