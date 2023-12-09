package main

import (
	"fmt"
	tictactoe "main/tic-tac-toe"
)

func main() {

	// we can take size as input from user
	size := 3
	fmt.Println("Enter board size:")
	_, err := fmt.Scan(&size)
	if err != nil {
		fmt.Println(err.Error())
	}
	board := tictactoe.Board{}
	board.BuildBoard(size)

	for true {
		id := board.GetCurrentPlayer().GetId()
		fmt.Printf("Player %d please make move.\n", id)

		var row, col int
		fmt.Println("Enter row:")
		_, err := fmt.Scan(&row)

		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Println("Enter column:")
		_, err = fmt.Scan(&col)

		if err != nil {
			fmt.Println(err.Error())
		}

		move := tictactoe.Move{
			Row:    row,
			Column: col,
			Player: board.GetCurrentPlayer(),
		}

		board.MakeMove(move)

		status := board.GetStatus()

		if status == 0 {
			continue
		} else if status == 1 {
			fmt.Println("player 1 won the game")
		} else if status == 1 {
			fmt.Println("player 2 won the game")
		}

		break
	}
}
