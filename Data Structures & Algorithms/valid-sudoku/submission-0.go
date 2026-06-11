func isValidSudoku(board [][]byte) bool {
	rows, columns, boxes := make([]map[byte]bool, 9), make([]map[byte]bool, 9), make([][]map[byte]bool, 3)
	for i := 0; i < 3; i++ {
		boxes[i] = make([]map[byte]bool, 3)
		for j := 0; j < 3; j++ {
			boxes[i][j] = make(map[byte]bool)
		}
	}
	for i := 0; i < 9; i++ {
		if rows[i] == nil {
			rows[i] = make(map[byte]bool)
		}
		for j := 0; j < 9; j++ {
			if columns[j] == nil {
				columns[j] = make(map[byte]bool)
			}
			n := board[i][j]
			if n == '.' {
				continue
			}
			if rows[i][n] {
				return false
			} else {
				rows[i][n] = true
			}
			if columns[j][n] {
				return false
			} else {
				columns[j][n] = true
			}
			if boxes[i/3][j/3][n] {
				return false
			} else {
				boxes[i/3][j/3][n] = true
			}
		}
	}
	return true
}
