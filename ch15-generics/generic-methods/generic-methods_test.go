package main

import (
	"fmt"
	"reflect"
	"slices"
	"testing"
)

func Test(t *testing.T) {
	welcome := email{"lane@example.com", "Welcome"}
	reminder := email{"dax@example.com", "Reminder"}
	in := inbox{events: []any{receipt{750}, welcome, receipt{200}, reminder, welcome}}
	type testCase struct {
		name string
		test func(*testing.T)
	}
	runCases := []testCase{
		{"select emails", func(t *testing.T) {
			testSelect(t, in, []email{welcome, reminder, welcome})
		}},
		{"select receipts from the same inbox", func(t *testing.T) {
			testSelect(t, in, []receipt{{750}, {200}})
		}},
	}
	submitCases := append(runCases, []testCase{
		{"no matching events", func(t *testing.T) {
			testSelect(t, in, []string{})
		}},
		{"empty inbox", func(t *testing.T) {
			testSelect(t, inbox{}, []email{})
		}},
		{"zero values and unrelated events", func(t *testing.T) {
			mixed := inbox{events: []any{nil, receipt{0}, 0, receipt{125}, "receipt"}}
			testSelect(t, mixed, []receipt{{0}, {125}})
		}},
		{"a new event type", func(t *testing.T) {
			type notification struct {
				subject string
				tags    []string
			}
			first := notification{"Update", []string{"account", "email"}}
			second := notification{"Alert", []string{"billing"}}
			mixed := inbox{events: []any{welcome, first, receipt{500}, second}}
			testSelect(t, mixed, []notification{first, second})
		}},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}
	passed, failed := 0, 0
	for i, tc := range testCases {
		fmt.Printf("---------------------------------\nSelectTest %d (%s):\n", i, tc.name)
		if t.Run(tc.name, tc.test) {
			fmt.Println("  Passed")
			passed++
		} else {
			fmt.Println("  Failed")
			failed++
		}
	}
	fmt.Println("---------------------------------")
	skipped := len(submitCases) - len(testCases)
	if skipped > 0 {
		fmt.Printf("Select: %d passed, %d failed, %d skipped\n", passed, failed, skipped)
	} else {
		fmt.Printf("Select: %d passed, %d failed\n", passed, failed)
	}
}

func testSelect[T any](t *testing.T, in inbox, want []T) {
	t.Helper()
	original := slices.Clone(in.events)
	var got []T = in.Select[T]()
	fmt.Printf(`  requested type: %T
  inbox:          %+v
  expected:       %+v
  actual:         %+v
`, *new(T), original, want, got)
	if !reflect.DeepEqual(in.events, original) {
		t.Errorf("inbox changed: expected %+v, actual %+v", original, in.events)
	}
	if len(got) != len(want) {
		t.Fatalf("expected %d events, actual %d", len(want), len(got))
	}
	for i := range want {
		if !reflect.DeepEqual(got[i], want[i]) {
			t.Errorf("event %d: expected %+v, actual %+v", i, want[i], got[i])
		}
	}
}

var withSubmit = true
