package main

type Analytics struct {
	MessagesTotal     int
	MessagesFailed    int
	MessagesSucceeded int
}

type Message struct {
	Recipient string
	Success   bool
}

//
func analyzeMessage(a *Analytics, msg Message) {
	a.MessagesTotal++
	if msg.Success {
		a.MessagesSucceeded++
	} else {
		a.MessagesFailed++
	}
}
