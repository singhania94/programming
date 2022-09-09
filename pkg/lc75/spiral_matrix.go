package main

import "fmt"

func SpiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	visited := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		visited[i] = make([]int, len(matrix[0]))
	}
	res := make([]int, len(matrix)*len(matrix[0]))
	so(matrix, visited, res, len(matrix), len(matrix[0]), 0, 0, 0, 0)
	return res
}

func so(matrix, visited [][]int, res []int, numRows, numCols, i, j, k, direction int) {
	fmt.Printf("now accessing index [%d,%d]\n", i, j)
	data := matrix[i][j]
	visited[i][j] = 1
	res[k] = data
	k++

	if direction == 0 {
		if j+1 < numCols && visited[i][j+1] == 0 {
			j++
		} else if i+1 < numRows && visited[i+1][j] == 0 {
			direction = 1
			i++
		} else {
			return
		}
	} else if direction == 1 {
		if i+1 < numRows && visited[i+1][j] == 0 {
			i++
		} else if j-1 >= 0 && visited[i][j-1] == 0 {
			direction = 2
			j--
		} else {
			return
		}
	} else if direction == 2 {
		if j-1 >= 0 && visited[i][j-1] == 0 {
			j--
		} else if i-1 >= 0 && visited[i-1][j] == 0 {
			direction = 3
			i--
		} else {
			return
		}
	} else if direction == 3 {
		if i-1 >= 0 && visited[i-1][j] == 0 {
			i--
		} else if j+1 < numCols && visited[i][j+1] == 0 {
			direction = 0
			j++
		} else {
			return
		}
	}

	so(matrix, visited, res, numRows, numCols, i, j, k, direction)
}
