package main

import "strings"

func removeProfanity(message *string) {
	// message is a string pointer
	if message == nil {
		return
	}
	msgCopy := *message
	msgCopy = strings.ReplaceAll(msgCopy, "fubb", "****")
	msgCopy = strings.ReplaceAll(msgCopy, "shiz", "****")
	msgCopy = strings.ReplaceAll(msgCopy, "witch", "*****")
	*message = msgCopy
}
