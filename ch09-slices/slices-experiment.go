package main

import "fmt"

func main() {
	mySlice := make([]int, 4)
	anotherSlice := []int{0, 1, 1, 2, 3, 5, 8}
	mySlice = append(mySlice, 0, 1, 2, 3)
	fmt.Println(mySlice)
	mySlice2 := mySlice[3:]
	fmt.Println(mySlice2)
	mySlice[3] = 10
	fmt.Println(mySlice)
	fmt.Println(mySlice2)
	fmt.Println(anotherSlice)
}
