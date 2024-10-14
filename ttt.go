package main

import (
	"math"
	"strings"
)

const (
	PlayerX = 'X'
	PlayerO = 'O'
	Empty   = ' '
)

func isEmpty(board [3][3]rune) bool {
	for _, row := range board {
		for _, cell := range row {
			if cell != Empty {
				return false
			}
		}
	}
	return true
}

func isFull(board [3][3]rune) bool {
	for _, row := range board {
		for _, cell := range row {
			if cell == Empty {
				return false
			}
		}
	}
	return true
}

func checkWinner(board [3][3]rune) rune {
	for i := 0; i < 3; i++ {
		if board[i][0] == board[i][1] && board[i][1] == board[i][2] && board[i][0] != Empty {
			return board[i][0]
		}
		if board[0][i] == board[1][i] && board[1][i] == board[2][i] && board[0][i] != Empty {
			return board[0][i]
		}
	}
	if board[0][0] == board[1][1] && board[1][1] == board[2][2] && board[0][0] != Empty {
		return board[0][0]
	}
	if board[0][2] == board[1][1] && board[1][1] == board[2][0] && board[0][2] != Empty {
		return board[0][2]
	}
	return Empty
}

// Function to get available moves (empty spots on the board)
func getAvailableMoves(board [3][3]rune) [][2]int {
	var moves [][2]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == Empty {
				moves = append(moves, [2]int{i, j})
			}
		}
	}
	return moves
}

func evaluate(board [3][3]rune, aiPlayer rune) int {
	winner := checkWinner(board)
	if winner == aiPlayer {
		return 1
	} else if winner != Empty && winner != aiPlayer {
		return -1
	}
	return 0
}

func minimaxAlphaBeta(board [3][3]rune, depth int, alpha int, beta int, isMaximizing bool, aiPlayer rune) int {
	score := evaluate(board, aiPlayer)

	if score == 1 || score == -1 {
		return score
	}
	if isFull(board) {
		return 0
	}

	var opponent rune
	if aiPlayer == PlayerX {
		opponent = PlayerO
	} else {
		opponent = PlayerX
	}

	if isMaximizing {
		best := math.MinInt32
		for _, move := range getAvailableMoves(board) {
			board[move[0]][move[1]] = aiPlayer
			best = max(best, minimaxAlphaBeta(board, depth+1, alpha, beta, false, aiPlayer))
			board[move[0]][move[1]] = Empty
			alpha = max(alpha, best)
			if beta <= alpha {
				break
			}
		}
		return best
	} else {
		best := math.MaxInt32
		for _, move := range getAvailableMoves(board) {
			board[move[0]][move[1]] = opponent
			best = min(best, minimaxAlphaBeta(board, depth+1, alpha, beta, true, aiPlayer))
			board[move[0]][move[1]] = Empty
			beta = min(beta, best)
			if beta <= alpha {
				break
			}
		}
		return best
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func findBestMove(board [3][3]rune, aiPlayer rune) (int, int) {
	bestScore := math.MinInt32
	moveRow, moveCol := -1, -1

	for _, move := range getAvailableMoves(board) {
		board[move[0]][move[1]] = aiPlayer
		moveScore := minimaxAlphaBeta(board, 0, math.MinInt32, math.MaxInt32, false, aiPlayer)
		board[move[0]][move[1]] = Empty

		if moveScore > bestScore {
			bestScore = moveScore
			moveRow, moveCol = move[0], move[1]
		}
	}
	return moveRow, moveCol
}

func parseBoard(movesString string) [3][3]rune {
	moves := strings.Split(movesString, "_")
	board := [3][3]rune{{Empty, Empty, Empty}, {Empty, Empty, Empty}, {Empty, Empty, Empty}}

	for _, move := range moves {
		split := strings.Split(move, "-")
		player := split[0]
		row, col := split[1], split[2]

		if player == "X" {
			board[int(row[0]-'0')][int(col[0]-'0')] = PlayerX
		} else if player == "O" {
			board[int(row[0]-'0')][int(col[0]-'0')] = PlayerO
		} else {
			board[int(row[0]-'0')][int(col[0]-'0')] = Empty
		}
	}

	return board
}
