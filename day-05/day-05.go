package main

import (
	"fmt"
)

func main() {
	test1 := readFile("./test-input.txt")
	fmt.Println("Solution Test 1", interpretLines(test1))

	input1 := readFile("./input.txt")
	fmt.Println("Solution 1", interpretLines(input1))

	test2 := readFile("./test-input.txt")
	fmt.Println("Solution Test 2", interpretAllLines(test2))

	input2 := readFile("./input.txt")
	fmt.Println("Solution 2", interpretAllLines(input2))

}

func interpretLines(input []string) int {
	moves := make([]*Move, len(input))
	for index, line := range input {
		moves[index] = NewMove(line)
	}

	field := NewField()
	for _, move := range moves {
		for _, touchedPoint := range move.touchedPoints() {
			field.increasePoint(touchedPoint)
		}
	}
	return field.highPoints()
}

func interpretAllLines(input []string) int {
	moves := make([]*Move, len(input))
	for index, line := range input {
		moves[index] = NewMove(line)
	}

	field := NewField()
	for _, move := range moves {
		for _, touchedPoint := range move.allTouchedPoints() {
			field.increasePoint(touchedPoint)
		}
	}
	return field.highPoints()
}
