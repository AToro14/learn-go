package main

import "strings"

func countDistinctWords(messages []string) int {
	distinctWords := make(map[string]struct{})
	for _, sentence := range messages {
		words := strings.Fields(sentence)
		for _, word := range words {
			// In this case the check below isn't necessary before inserting
			// because we are only storing an empty struct
			if _, ok := distinctWords[strings.ToLower(word)]; !ok {
				distinctWords[strings.ToLower(word)] = struct{}{}
			}
		}
	}
	return len(distinctWords)
}
