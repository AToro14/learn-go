package main

import "strings"

func countDistinctWords(messages []string) int {
	distinctWords := make(map[string]struct{})
	for _, sentence := range messages {
		words := strings.Fields(strings.ToLower(sentence))
		for _, word := range words {
			distinctWords[word] = struct{}{}
		}
	}
	return len(distinctWords)
}
