package main

import "unicode"

func isValidPassword(password string) bool {
	rPassword := []rune(password)
	hasUpper := false
	hasDigit := false
	pLen := len(rPassword)
	if pLen < 5 || pLen > 12 {
		return false
	}
	for i := 0; i < pLen; i++ {
		if unicode.IsDigit(rPassword[i]) {
			hasDigit = true
		}
		if unicode.IsUpper(rPassword[i]) {
			hasUpper = true
		}
		if hasUpper && hasDigit {
			return true
		}
	}
	return false
}
