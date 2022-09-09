package main

func FloodFill(image [][]int, sr int, sc int, color int) [][]int {
	rows := len(image)
	cols := len(image[0])
	if image[sr][sc] != color {
		fill(image, rows, cols, sr, sc, image[sr][sc], color)
	}
	return image
}

func fill(image [][]int, rows, cols, sr, sc, oldColor, color int) {
	if image[sr][sc] != oldColor {
		return
	}

	image[sr][sc] = color
	if sr-1 >= 0 {
		fill(image, rows, cols, sr-1, sc, oldColor, color)
	}
	if sr+1 < rows {
		fill(image, rows, cols, sr+1, sc, oldColor, color)
	}
	if sc-1 >= 0 {
		fill(image, rows, cols, sr, sc-1, oldColor, color)
	}
	if sc+1 < cols {
		fill(image, rows, cols, sr, sc+1, oldColor, color)
	}
}
