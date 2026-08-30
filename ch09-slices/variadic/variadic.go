package main

func sum(nums ...int) int {
	_sum := 0
	for i := 0; i < len(nums); i++ {
		_sum += nums[i]
	}
	return _sum
}
