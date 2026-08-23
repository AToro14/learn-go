package main

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	originalMessages := [3]string{primary, secondary, tertiary}
	lenPrimary := len(primary)
	lenSecondary := len(secondary)
	lenTertiary := len(tertiary)
	costOfEachMessage := [3]int{lenPrimary, lenPrimary + lenSecondary, lenPrimary + lenSecondary + lenTertiary}
	return originalMessages, costOfEachMessage
}
