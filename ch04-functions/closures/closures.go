package main

func adder() func(int) int {
	// The scope of `sum` is the _whole_ adder() function
	// A var declared inside the return func() would be inaccessable from the
	// block level `sum` is accessible from
	sum := 0
	return func(input int) int {
		sum += input
		return sum
	}
}
