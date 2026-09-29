package main

import "fmt"
// type emailEnvelope struct {
//	recipient string
//	payload   string
// }
//
// type paymentEnvelope struct {
//	recipient string
//	payload   int
// }

type envelope[T any] struct {
	recipient string
	payload   T
}

func createEnvelopes[T any](recipients []string, payload T) []envelope[T] {
	rLen := len(recipients)
	// envelopes := make([]envelope[T], rLen)
	envelopes := []envelope[T]{}
	newRecipients := make(map[string]struct{})
	fmt.Printf("envelopes: %v\n", envelopes)
	fmt.Printf("rLen: %d\n", rLen)
	if rLen == 0 {
		return envelopes
	}
	for _, recipient := range recipients {
		fmt.Printf("recipient attempt: %s\n", recipient)
		// if _, ok := newRecipients[recipient]; ok {
		//	fmt.Printf("recipient %s already found\n", recipient)
		//	continue
		// }
		newRecipients[recipient] = struct{}{}
		fmt.Println(newRecipients)
		envelopes = append(envelopes, envelope[T]{recipient: recipient, payload: payload})
	}
	fmt.Printf("envelopes: %v\n", envelopes)
	return envelopes
}

