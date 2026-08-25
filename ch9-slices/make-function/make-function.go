package main

func getMessageCosts(messages []string) []float64 {
	messageCosts := make([]float64, len(messages))
	for i := 0; i < len(messageCosts); i++ {
		messageCosts[i] = float64(len(messages[i])) * 0.01
	}
	return messageCosts
}
