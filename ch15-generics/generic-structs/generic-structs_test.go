package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		name string
		test func(*testing.T)
	}
	runCases := []testCase{
		{"email body", func(t *testing.T) {
			testEnvelopes(t, []string{"lane@example.com", "dax@example.com"}, "Welcome to Textio!")
		}},
		{"payment amount", func(t *testing.T) {
			testEnvelopes(t, []string{"billing@example.com"}, 750)
		}},
	}
	submitCases := append(runCases, []testCase{
		{"no recipients", func(t *testing.T) {
			testEnvelopes(t, []string{}, "Welcome to Textio!")
		}},
		{"duplicate recipients", func(t *testing.T) {
			testEnvelopes(t, []string{"billing@example.com", "billing@example.com"}, 0)
		}},
		{"structured notification", func(t *testing.T) {
			type notification struct {
				subject string
				tags    []string
			}
			testEnvelopes(t, []string{"allan@example.com", "lane@example.com", "dax@example.com"},
				notification{"Updated receipt", []string{"payment", "corrected"}},
			)
		}},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}
	passed, failed := 0, 0
	for i, tc := range testCases {
		fmt.Printf("---------------------------------\nEnvelopeTest %d (%s):\n", i, tc.name)
		if t.Run(tc.name, tc.test) {
			passed++
		} else {
			failed++
		}
	}

	fmt.Println("---------------------------------")
	skipped := len(submitCases) - len(testCases)
	if skipped > 0 {
		fmt.Printf("Envelopes: %d passed, %d failed, %d skipped\n", passed, failed, skipped)
	} else {
		fmt.Printf("Envelopes: %d passed, %d failed\n", passed, failed)
	}
}

func testEnvelopes[T any](t *testing.T, recipients []string, payload T) {
	t.Helper()
	var envelopes []envelope[T] = createEnvelopes(recipients, payload)
	fmt.Printf(`  recipients:       %q
  payload:          %+v
  expected count:   %d
  actual count:     %d
`, recipients, payload, len(recipients), len(envelopes))
	if len(envelopes) != len(recipients) {
		t.Fatal("envelope count does not match")
	}
	for i, e := range envelopes {
		var stored T = e.payload
		fmt.Printf(`  envelope %d:
    expected recipient: %q
    actual recipient:   %q
    expected payload:   %+v
    actual payload:     %+v
`, i, recipients[i], e.recipient, payload, stored)
		if e.recipient != recipients[i] || !reflect.DeepEqual(stored, payload) {
			t.Errorf("envelope %d does not match", i)
		}
	}
}

var withSubmit = true
