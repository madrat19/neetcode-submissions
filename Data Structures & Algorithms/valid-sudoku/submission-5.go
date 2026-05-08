func isValidSudoku(board [][]byte) bool {
	rows := [9][9]bool{}
	cols := [9][9]bool{}
	boxs := [9][9]bool{}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}

			num := int(board[i][j] - '1')

			if rows[i][num] == true {
				return false
			}
			rows[i][num] = true

			if cols[j][num] == true {
				return false
			}
			cols[j][num] = true

			if boxs[(i/3)*3+j/3][num] == true {
				return false
			}
			boxs[(i/3)*3+j/3][num] = true
		}
	}
	return true
}
