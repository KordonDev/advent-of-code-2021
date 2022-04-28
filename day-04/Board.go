package main

import "log"

type Board struct {
	rows    [][]*Field
	columns [][]*Field
	won     bool
}

type Field struct {
	number   int
	selected bool
}

func (board *Board) mark(number int) {
	for _, row := range board.rows {
		for _, field := range row {
			if field.number == number {
				field.selected = true
			}
		}
	}

	for _, column := range board.columns {
		for _, field := range column {
			if field.number == number {
				field.selected = true
			}
		}
	}
}

func (board *Board) bingo() bool {
	for _, row := range board.rows {
		if board.won {
			return true
		}
		bingoRow := true
		for _, field := range row {
			if !field.selected {
				bingoRow = false
			}
		}
		if bingoRow {
			board.won = true
			return true
		}
	}

	for _, column := range board.columns {
		bingoColumn := true
		for _, field := range column {
			if !field.selected {
				bingoColumn = false
			}
		}
		if bingoColumn {
			board.won = true
			return true
		}
	}

	return false
}

func (board *Board) calculateResult(winningNumber int) int {
	sum := 0
	for _, row := range board.rows {
		for _, field := range row {
			if !field.selected {
				sum = sum + field.number
			}
		}
	}

	return sum * winningNumber
}

func (board *Board) print() {
	log.Println("Board ", board.won)
	for _, row := range board.rows {
		for _, field := range row {
			log.Print("{", field.number, " ", field.selected, "} ")
		}
		log.Println()
	}
	log.Println(" -- ")
	for _, column := range board.columns {
		for _, field := range column {
			log.Print("{", field.number, " ", field.selected, "} ")
		}
		log.Println()
	}
	log.Println(" -- -- -- ")
}
