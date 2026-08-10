package main

import "fmt"

func printReports(intro, body, outro string) {
	printCostReport(func(x string) int {
		return 2 * len(x)
	}, intro)
	printCostReport(func(y string) int {
		return 3 * len(y)
	}, body)
	printCostReport(func(z string) int {
		return 4 * len(z)
	}, outro)
}

// don't touch below this line

func main() {
	printReports(
		"Welcome to the Hotel California",
		"Such a lovely place",
		"Plenty of room at the Hotel California",
	)
}

func printCostReport(costCalculator func(string) int, message string) {
	cost := costCalculator(message)
	fmt.Printf(`Message: "%s" Cost: %v cents`, message, cost)
	fmt.Println()
}
