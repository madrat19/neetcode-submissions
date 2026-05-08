func isValidSudoku(board [][]byte) bool {
	//strings
	for i := 0; i < 9; i++ {
		counter := make(map[byte]int, 9)
		for j := 0; j < 9; j++ {
			elem := board[i][j]
			if elem != '.' {
				counter[elem]++
				if counter[elem] == 2 {
					return false
				}
			}
		}
	}

	//columns
	for j := 0; j < 9; j++ {
		counter := make(map[byte]int, 9)
		for i := 0; i < 9; i++ {
			elem := board[i][j]
			if elem != '.' {
				counter[elem]++
				if counter[elem] == 2 {
					return false
				}
			}
		}
	}

	//squares
	coordinates := []int{0, 3, 6}
	for _, i := range coordinates {
		for _, j := range coordinates {
			counter := make(map[byte]int)
			for x := i; x < i+3; x++ {
				for y := j; y < j+3; y++ {
					elem := board[x][y]
					if elem != '.' {
						counter[elem]++
						if counter[elem] == 2 {
							return false
						}
					}
				}
			}

		}
	}
	return true
}
