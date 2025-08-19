package main

func validateRows(board [][]byte) bool {
	for i := 0; i < len(board); i++ {
		flag := map[byte]bool{
			'1': false,
			'2': false,
			'3': false,
			'4': false,
			'5': false,
			'6': false,
			'7': false,
			'8': false,
			'9': false,
		} // Refresh flag for retrieving each row
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == '.' {
				continue
			}
			if v, ok := flag[board[i][j]]; ok && v == true {
				return false
			}
			flag[board[i][j]] = true
		}
	}
	return true
}

func validateColumns(board [][]byte) bool {
	for i := 0; i < len(board); i++ {
		flag := map[byte]bool{
			'1': false,
			'2': false,
			'3': false,
			'4': false,
			'5': false,
			'6': false,
			'7': false,
			'8': false,
			'9': false,
		} // Refresh flag for retrieving each column
		for j := 0; j < len(board[i]); j++ {
			if board[j][i] == '.' {
				continue
			}
			if v, ok := flag[board[j][i]]; ok && v == true {
				return false
			}
			flag[board[j][i]] = true
		}
	}
	return true
}

func validateSubBoxes(board [][]byte, xgrid, ygrid int) bool {
	flag := map[byte]bool{
		'1': false,
		'2': false,
		'3': false,
		'4': false,
		'5': false,
		'6': false,
		'7': false,
		'8': false,
		'9': false,
	} // Refresh flag for retrieving each sub-box
	for i := 3 * xgrid; i < 3*xgrid+3; i++ {
		for j := 3 * ygrid; j < 3*ygrid+3; j++ {
			if board[i][j] == '.' {
				continue
			}
			if v, ok := flag[board[i][j]]; ok && v == true {
				return false
			}
			flag[board[i][j]] = true
		}
	}
	return true
}

func isValidSudoku(board [][]byte) bool {

	var isValidSubBoxes, isValidRows, isValidColumns bool
	// Processing sub-boxes
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			isValidSubBoxes = validateSubBoxes(board, i, j)
			if isValidSubBoxes == false {
				return false
			}
		}
	}

	// Processing rows
	if isValidRows = validateRows(board); isValidRows == false {
		return false
	}

	// Processing columns
	if isValidColumns = validateColumns(board); isValidColumns == false {
		return false
	}

	return true
}

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	println(isValidSudoku(board)) // Output: true

	board2 := [][]byte{
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	println(isValidSudoku(board2)) // Output: false

	board3 := [][]byte{
		{'.', '.', '.', '.', '5', '.', '.', '1', '.'},
		{'.', '4', '.', '3', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '3', '.', '.', '1'},
		{'8', '.', '.', '.', '.', '.', '.', '2', '.'},
		{'.', '.', '2', '.', '7', '.', '.', '.', '.'},
		{'.', '1', '5', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '2', '.', '.', '.'},
		{'.', '2', '.', '9', '.', '.', '.', '.', '.'},
		{'.', '.', '4', '.', '.', '.', '.', '.', '.'},
	}
	println(isValidSudoku(board3)) // Output: false

}
