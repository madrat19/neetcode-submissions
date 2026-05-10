func isValidSudoku(board [][]byte) bool {
	rows := [9]int{}
	cols := [9]int{}
	boxs := [9]int{}

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			value := board[r][c]
			if value == '.' {
				continue
			}

			num := value - '1'
			bit := 1 << num
			box := (r/3)*3 + c/3

			if rows[r]&bit != 0 ||
				cols[c]&bit != 0 ||
				boxs[box]&bit != 0 {
				return false
			}

			rows[r] |= bit
			cols[c] |= bit
			boxs[box] |= bit
		}
	}
	return true
}
