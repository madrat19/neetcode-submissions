func isValidSudoku(board [][]byte) bool {
	rows := make(map[int]map[byte]int, 9)
	columns := make(map[int]map[byte]int, 9)
	squares := make(map[[2]int]map[byte]int, 9)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			if rows[i] == nil {
				rows[i] = make(map[byte]int, 9)
			}
			if columns[j] == nil {
				columns[j] = make(map[byte]int, 9)
			}
			if squares[[2]int{i / 3, j / 3}] == nil {
				squares[[2]int{i / 3, j / 3}] = make(map[byte]int, 9)
			}
			rows[i][board[i][j]]++
			columns[j][board[i][j]]++
			squares[[2]int{i / 3, j / 3}][board[i][j]]++

			if rows[i][board[i][j]] == 2 || columns[j][board[i][j]] == 2 || squares[[2]int{i / 3, j / 3}][board[i][j]] == 2 {
				return false
			}
		}
	}
	return true
}
