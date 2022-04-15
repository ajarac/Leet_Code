package valid_sudoku

func isValidSudoku(board [][]byte) bool {

	for i := 0; i < 9; i++ {
		setRows := make(map[byte]bool)
		for r := 0; r < 9; r++ {
			if board[i][r] == byte('.') {
				continue
			}
			if _, ok := setRows[board[i][r]]; ok == true {
				return false
			} else {
				setRows[board[i][r]] = true
			}
		}

		setColumns := make(map[byte]bool)
		for c := 0; c < 9; c++ {
			if board[c][i] == byte('.') {
				continue
			}
			if _, ok := setColumns[board[c][i]]; ok == true {
				return false
			} else {
				setColumns[board[c][i]] = true
			}
		}

		setSquare := make(map[byte]bool)
		for j := 0; j < 3; j++ {
			row := j + (i/3)*3
			for z := 0; z < 3; z++ {
				column := z + (i%3)*3
				if board[row][column] == byte('.') {
					continue
				}
				if _, ok := setSquare[board[row][column]]; ok == true {
					return false
				} else {
					setSquare[board[row][column]] = true
				}
			}
		}
	}

	return true
}
