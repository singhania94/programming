package main

func findBall(grid [][]int) []int {
	numRows, numCols := len(grid), len(grid[0])
	res := make([]int, numCols)
	for k := 0; k < numCols; k++ {
		i, j := 0, k
		direction := 1

		for i < numRows && j >= 0 && j <= numCols {
			if direction == 0 {
				if grid[i][j] == 1 {
					direction = 1
					i++
					continue
				}
				break
			} else if direction == 1 {
				if grid[i][j] == 1 {
					direction = 0
					j++
					continue
				}
				direction = 2
				j--
				continue
			}
			if grid[i][j] == -1 {
				direction = 1
				i++
				continue
			}
			break
		}

		if i == numRows {
			res[k] = j
		}
		res[k] = -1
	}
	return res
}

func FindBallRecursion(grid [][]int) []int {
	numRows, numCols := len(grid), len(grid[0])
	res := make([]int, numCols)
	for j := 0; j < numCols; j++ {
		res[j] = next(grid, 1, numRows, numCols, 0, j)
	}
	return res
}

func next(grid [][]int, direction, numRows, numCols, i, j int) int {
	if i == numRows {
		return j
	} else if j == -1 || j == numCols {
		return -1
	}

	if direction == 0 {
		if grid[i][j] == 1 {
			return next(grid, 1, numRows, numCols, i+1, j)
		}
		return -1
	} else if direction == 1 {
		if grid[i][j] == 1 {
			return next(grid, 0, numRows, numCols, i, j+1)
		}
		return next(grid, 2, numRows, numCols, i, j-1)
	}
	if grid[i][j] == 1 {
		return -1
	}
	return next(grid, 1, numRows, numCols, i+1, j)
}
