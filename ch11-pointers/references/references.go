package main

import "strings"

func removeProfanity(message *string) {
	// message is a string pointer
	*message = strings.ReplaceAll(*message, "fubb", "****")
	*message = strings.ReplaceAll(*message, "shiz", "****")
	*message = strings.ReplaceAll(*message, "witch", "*****")

}
