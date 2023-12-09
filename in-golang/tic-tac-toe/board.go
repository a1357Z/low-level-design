package tictactoe

import "fmt"

type Board struct {
	size          int
	grid          [][]int
	currentPlayer Player
	otherPlayer   Player
}

// exported methods
func (b *Board) GetCurrentPlayer() *Player {
	return &b.currentPlayer
}

func (b *Board) MakeMove(move Move) {
	valid := b.validateMove(move)

	if !valid {
		return
	}

	row := move.Row - 1
	col := move.Column - 1

	b.grid[row][col] = b.currentPlayer.id

	// toggle the players
	opponent := b.currentPlayer
	b.currentPlayer = b.otherPlayer
	b.otherPlayer = opponent

	b.printBoard()
}

// 0 --> nothing happened, 1 --> player 1 won, 2 --> player 2 won
func (b *Board) GetStatus() int {

	for r := 0; r < b.size; r++ {
		p1count, p2count := 0, 0
		for c := 0; c < b.size; c++ {
			if b.grid[r][c] == 1 {
				p1count++
			} else if b.grid[r][c] == 2 {
				p2count++
			}

			if p1count == b.size {
				return 1
			} else if p2count == b.size {
				return 2
			}
		}
	}

	for c := 0; c < b.size; c++ {
		p1count, p2count := 0, 0
		for r := 0; r < b.size; r++ {
			if b.grid[r][c] == 1 {
				p1count++
			} else if b.grid[r][c] == 2 {
				p2count++
			}

			if p1count == b.size {
				return 1
			} else if p2count == b.size {
				return 2
			}
		}
	}

	// 1st diagonal
	p1count, p2count := 0, 0
	for r := 0; r < b.size; r++ {
		if b.grid[r][r] == 1 {
			p1count++
		} else if b.grid[r][r] == 2 {
			p2count++
		}
	}

	if p1count == b.size {
		return 1
	} else if p2count == b.size {
		return 2
	}

	// second diagonal
	p1count = 0
	p2count = 0
	for r := 0; r < b.size; r++ {
		if b.grid[r][b.size-r-1] == 1 {
			p1count++
		} else if b.grid[r][b.size-r-1] == 2 {
			p2count++
		}
	}

	if p1count == b.size {
		return 1
	} else if p2count == b.size {
		return 2
	}

	return 0
}

func (b *Board) BuildBoard(size int) {
	grid := make([][]int, size)

	for i := 0; i < size; i++ {
		row := []int{}
		for j := 0; j < size; j++ {
			row = append(row, 0)
		}
		grid[i] = row
	}

	// var grid [size][size]int = [size][size]int{0}

	b.grid = grid

	player1 := Player{
		id: 1,
	}

	player2 := Player{
		id: 2,
	}

	b.currentPlayer = player1
	b.otherPlayer = player2
	b.size = size

	fmt.Printf("Board created of size: %d\n", size)
	fmt.Printf("You can enter row: 1...%d\n", size)
	fmt.Printf("You can enter column: 1...%d\n", size)
}

// internal methods
func (b *Board) validateMove(move Move) bool {
	row := move.Row - 1
	col := move.Column - 1

	// invalid cell selected
	if row < 0 || row >= b.size || col < 0 || col >= b.size {
		fmt.Println("invalid cell selected")
		return false
	}

	// this cell is already occupied
	if b.grid[row][col] != 0 {
		fmt.Println("this cell is already occupied")
		return false
	}

	return true
}

func (b *Board) printBoard() {
	for r := 0; r < b.size; r++ {
		for c := 0; c < b.size; c++ {
			fmt.Print(b.grid[r][c])
		}
		fmt.Println()
	}
}
