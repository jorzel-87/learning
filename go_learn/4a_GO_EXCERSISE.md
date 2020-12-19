 Implement WordCount. It should return a map of the counts of each ?word? in the string s. The wc.Test function runs a test suite against the provided function and prints success or failure.

You might find strings.Fields helpful. 

package main

import (
	"golang.org/x/tour/wc"
	"strings"
	"fmt"
)

func WordCount(s string) map[string]int {
	wordMap := make(map[string]int)
	words := strings.Fields(s)
	fmt.Println(s)
	for _, word := range words {
			wordMap[word]++
	}
	return wordMap
}


func main() {
	wc.Test(WordCount)
}
