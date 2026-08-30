package main

func getMaxMessagesToSend(costMultiplier float64, budgetInPennies int) int {
	actualCostInPennies := 1.0
	maxMessagesToSend := 1
	balance := float64(budgetInPennies) - actualCostInPennies
	for balance > 0 {
		actualCostInPennies *= costMultiplier
		balance -= actualCostInPennies
		maxMessagesToSend++
	}
	if balance < 0 {
		maxMessagesToSend--
	}
	return maxMessagesToSend
}
