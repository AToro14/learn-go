package main

func createMatrix(rows, cols int) [][]int {
	matrix := [][]int{}
	for i := 0; i < rows; i++ {
		row := []int{}
		for j := 0; j < cols; j++ {
			row = append(row, i*j)
		}
		matrix = append(matrix, row)
	}
	return matrix
}
