package main

// takes a message (string) and a function that takes a string which returns a string
func reformat(message string, formatter func(string) string) string {
	// message = formatter(message)
	// message = formatter(message)
	// message = formatter(message)
	message = formatter(formatter(formatter(message)))
	return "TEXTIO: " + message
}
