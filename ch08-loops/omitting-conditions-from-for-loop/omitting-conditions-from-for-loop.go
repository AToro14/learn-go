package main

func maxMessages(thresh int) int {
	messages := 0
	for i := 0; ; i++ {
		thresh -= 100 + i
		if thresh >= 0 {
			messages++
		} else {
			return messages
		}
	}
}
