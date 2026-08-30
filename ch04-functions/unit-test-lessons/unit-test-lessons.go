package uTest

func getMonthlyPrice(tier string) int {
	var pennies int
	switch tier {
		case "basic":
			pennies = 10000
		case "premium":
			pennies = 15000
		case "enterprise":
			pennies = 50000
		default:
			pennies = 0
	}
	// eh
	return pennies
}
