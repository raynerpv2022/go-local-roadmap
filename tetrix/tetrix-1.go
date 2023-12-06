package main

import (
	"fmt"
)

func printBoard(board [][]string) {
	fmt.Println()
	defer fmt.Println()
	for i := 0; i < 10; i++ {
		fmt.Println()
		for j := 0; j < 10; j++ {
			fmt.Print(board[i][j])
		}
	}

}

func movDown(board [][]string) [][]string {
	newBoard := newBoard()
	for indexRow, row := range board {
		for indexCol, col := range row {
			if col == "🔳" {
				if indexRow < len(row)-1 {
					newBoard[indexRow+1][indexCol] = "🔳"
				} else {
					return board
				}

			}

		}
	}
	return newBoard
}

func movLeft(board [][]string) [][]string {
	newBoard := newBoard()
	for indexRow, row := range board {
		for indexCol, col := range row {
			if col == "🔳" {
				if indexCol == 0 {
					return board
				}
				newBoard[indexRow][indexCol-1] = "🔳"
			}
		}
	}

	return newBoard
}
func movRight(board [][]string) [][]string {
	newBoard := newBoard()
	for indexrow, row := range board {
		for indexCol, col := range row {
			if col == "🔳" {

				if indexCol < len(board[indexrow])-1 {
					newBoard[indexrow][indexCol+1] = "🔳"

				} else {
					return board
				}

			}
		}
	}
	return newBoard

}

func GetKeyValue(board [][]string, key int) [][]string {

	switch key {
	case 1:
		board = movDown(board)
	case 2:
		board = movLeft(board)
	case 3:
		board = movRight(board)

	}
	return board
}

func getInput() string {
	var input string
	fmt.Println("Pess (A) for left, (D) for right, (X) for down")
	fmt.Scanf("%s", &input)

	return input

}

func getKey(board [][]string) {
	// printBoard(board)
	for {
		keyinput := getInput()
		switch keyinput {
		case "x":
			board = movDown(board)
		case "a":
			board = movLeft(board)
		case "d":
			board = movRight(board)

		}
		printBoard(board)
	}
}
func newBoard() [][]string {
	return [][]string{
		{"🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲"},
		{"🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲"},
		{"🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲"},
		{"🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲"},
		{"🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲"},
		{"🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲"},
		{"🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲"},
		{"🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲"},
		{"🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲"},
		{"🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲"},
	}
}

// func cantItmoveDown(board [][]string) bool {
// 	canMove := true
// 	for indexRow, row := range board {

// 	}

// 	return canMove
// }

func main() {
	fmt.Println("============== T E T R I X =====================")
	// black := "🔳"
	// white := "🔲"  aa
	board := [][]string{
		{"🔳", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲"},
		{"🔳", "🔳", "🔳", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲"},
		{"🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲"},
		{"🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲"},
		{"🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲"},
		{"🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲"},
		{"🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲"},
		{"🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲"},
		{"🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲"},
		{"🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲", "🔲"},
	}

	// printBoard(board)
	// board = GetKeyValue(board, 3)
	// printBoard(board)
	// board = GetKeyValue(board, 3)
	// printBoard(board)
	// board = GetKeyValue(board, 3)
	// printBoard(board)
	// board = GetKeyValue(board, 3)
	// printBoard(board)
	// board = GetKeyValue(board, 3)
	// printBoard(board)
	// board = GetKeyValue(board, 3)
	// printBoard(board)
	// board = GetKeyValue(board, 3)
	// printBoard(board)
	// board = GetKeyValue(board, 3)
	// printBoard(board)

	// board = GetKeyValue(board, 3)
	// printBoard(board)
	// board = GetKeyValue(board, 3)
	// printBoard(board)
	// board = GetKeyValue(board, 3)
	// printBoard(board)
	// board = GetKeyValue(board, 3)
	// printBoard(board)
	// board = GetKeyValue(board, 3)
	// printBoard(board)
	// board = GetKeyValue(board, 3)
	// printBoard(board)

	// getKey(board)
	printBoard(board)
	getKey(board)

}
