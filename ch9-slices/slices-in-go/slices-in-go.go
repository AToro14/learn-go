package main

import (
	"errors"
)

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {
	var e []string
	if plan == planPro {
		return messages[:], nil
	} else if plan == planFree {
		return messages[:2], nil
	} else {
		return e, errors.New("unsupported plan")
	}
}
