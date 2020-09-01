package array

import "math"

func setZeroes(matrix [][]int) {
	//SetZeros_1(matrix)
	//SetZeros_2(matrix)
	SetZeros(matrix)
}

func SetZeros_1(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	lines := make(map[int]bool)
	rows := make(map[int]bool)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				lines[i] = true
				rows[j] = true
			}
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if lines[i] || rows[j] {
				matrix[i][j] = 0
			}
		}
	}
}

func SetZeros_2(matrix [][]int) {
	const placeholder = -math.MaxInt32
	if len(matrix) == 0 {
		return
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				for jj := 0; jj < len(matrix[0]); jj++ {
					if matrix[i][jj] != 0 {
						matrix[i][jj] = placeholder
					}
				}
				for ii := 0; ii < len(matrix); ii++ {
					if matrix[ii][j] != 0 {
						matrix[ii][j] = placeholder
					}
				}
			}
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == placeholder {
				matrix[i][j] = 0
			}
		}
	}
}

func SetZeros(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	var firstRow, firstColumn bool
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				if i == 0 {
					firstRow = true
				}
				if j == 0 {
					firstColumn = true
				}
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := 1; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			for j := 1; j < len(matrix[0]); j++ {
				matrix[i][j] = 0
			}
		}
	}
	for j := 1; j < len(matrix[0]); j++ {
		if matrix[0][j] == 0 {
			for i := 1; i < len(matrix); i++ {
				matrix[i][j] = 0
			}
		}
	}
	if firstRow {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[0][j] = 0
		}
	}
	if firstColumn {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}
}
