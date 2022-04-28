package main

import (
	"fmt"
	"strings"
)

func main() {
	test1 := readFile("./test-input.txt")
	pulledNumbersTest := stringsToInt(strings.Split(test1[0], ","))
	boards := *getFieldsForInput(test1[2:])

	fmt.Println("Solution Test 1", playToWin(pulledNumbersTest, &boards))

	input1 := readFile("./input.txt")
	pulledNumbers := stringsToInt(strings.Split(input1[0], ","))
	boards = *getFieldsForInput(input1[2:])
	fmt.Println("Solution 1", playToWin(pulledNumbers, &boards))

	test2 := readFile("./test-input.txt")
	boards = *getFieldsForInput(test2[2:])
	fmt.Println("Solution Test 2", playToLose(pulledNumbersTest, &boards))

	input2 := readFile("./input.txt")
	boards = *getFieldsForInput(input2[2:])
	fmt.Println("Solution 2", playToLose(pulledNumbers, &boards))
}

func playToWin(pulledNumbers []int, boards *[]Board) int {
	for _, number := range pulledNumbers {
		for _, board := range *boards {
			board.mark(number)

			if board.bingo() {
				return board.calculateResult(number)
			}
		}
	}
	return 0
}

func playToLose(pulledNumbers []int, boards *[]Board) int {
	var wonBoardsIndex []int
	for _, number := range pulledNumbers {
		for index, board := range *boards {
			board.mark(number)
			if board.bingo() {
				contains := false
				for _, wonIndex := range wonBoardsIndex {
					if wonIndex == index {
						contains = true
					}
				}
				if !contains {
					wonBoardsIndex = append(wonBoardsIndex, index)
				}
			}

		}
		if len(wonBoardsIndex) == len(*boards) {
			return (*boards)[wonBoardsIndex[len(wonBoardsIndex)-1]].calculateResult(number)
		}
	}
	return 0
}

func getFieldsForInput(input []string) *[]Board {
	var boards []Board
	currentBoard := Board{won: false}
	for _, row := range input {
		if row == "" {
			boards = append(boards, currentBoard)
			currentBoard = Board{won: false}
			continue
		}

		currentRowAsString := strings.Split(row, " ")
		var currentRow []*Field
		for _, number := range currentRowAsString {
			if number != "" {
				currentRow = append(currentRow, &Field{number: stringToInt(number), selected: false})
			}
		}
		currentBoard.rows = append(currentBoard.rows, currentRow)

		if currentBoard.columns == nil {
			currentBoard.columns = make([][]*Field, len(currentRow))
		}

		for index, numberField := range currentRow {
			currentBoard.columns[index] = append(currentBoard.columns[index], &Field{number: numberField.number, selected: false})
		}
	}
	boards = append(boards, currentBoard)
	return &boards
}
