package main

type envelope[T any] struct {
	recipient string
	payload   T
}

func createEnvelopes[T any](recipients []string, payload T) []envelope[T] {
	envelopes := make([]envelope[T], len(recipients))
	for i, recipient := range recipients {
		envelopes[i] = envelope[T]{
			recipient: recipient,
			payload:   payload,
		}
	}
	return envelopes
}
