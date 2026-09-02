package main

import "unicode/utf8"

func getNameCounts(names []string) map[rune]map[string]int {
	nameCounts := make(map[rune]map[string]int)
	for _, name := range names {
		letter, _ := utf8.DecodeRuneInString(name)
		if _, ok := nameCounts[letter]; !ok {
			subMap := map[string]int{
				name: 0}
			nameCounts[letter] = subMap
		}
		nameCounts[letter][name]++
	}
	return nameCounts
}
