package controllers

import (
	"fmt"
	"os"

	"github.com/gregoflash05/terminal-snake-game/models"
)

func RequestBoardInfo() (height, width int) {
	fmt.Print("Enter board size information \n")

	fmt.Print("Enter height: ")
	_, err := fmt.Scanln(&width)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	fmt.Print("Enter width: ")
	_, err = fmt.Scanln(&height)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	return
}

func GetnewBoard(width, height int) (board models.Board) {
	cells := make([][]int, height)

	for i := range cells {
		cells[i] = make([]int, width)
	}

	board = models.Board{
		Cells:  cells,
		Width:  width,
		Height: height,
	}

	return
}
