package main

func indexOfFirstBadWord(msg []string, badWords []string) int {
	for i, word := range msg {
		for _, bWord := range badWords {
			if word == bWord {
				return i
			}
		}
	}
	return -1
}
