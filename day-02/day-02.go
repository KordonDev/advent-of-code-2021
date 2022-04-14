package main

import (
	"fmt"
	"strings"
)

type Position struct {
	horizontal, depth, aim int
	direction              string
}

func (p *Position) add(move Position) {
	p.horizontal = p.horizontal + move.horizontal
	p.depth = p.depth + move.depth
}

func (p *Position) add2(move Position) {
	if move.direction == "forward" {
		p.horizontal = p.horizontal + move.horizontal
		p.depth = p.depth + p.aim*move.aim
	} else {
		p.aim = p.aim + move.aim
	}
}

func main() {
	test1 := readFile("./test-input.txt")
	subTest := moveSubmarine(inputToMovements(test1))
	fmt.Println("Solution Test 1", subTest.horizontal*subTest.depth)

	input1 := readFile("./input.txt")
	sub1 := moveSubmarine(inputToMovements(input1))
	fmt.Println("Solution 1", sub1.horizontal*sub1.depth)

	test2 := readFile("./test-input.txt")
	subTest2 := moveSubmarine2(inputToMovements(test2))
	fmt.Println("Solution Test 2", subTest2.horizontal*subTest2.depth)

	input2 := readFile("./input.txt")
	sub2 := moveSubmarine2(inputToMovements(input2))
	fmt.Println("Solution 2", sub2.horizontal*sub2.depth)
}

func inputToMovements(lines []string) []Position {
	var positions []Position
	for _, line := range lines {
		s := strings.Split(line, " ")
		direction, move := s[0], stringToInt(s[1])
		switch direction {
		case "up":
			positions = append(positions, Position{0, move * -1, move * -1, direction})
		case "down":
			positions = append(positions, Position{0, move, move, direction})
		case "forward":
			positions = append(positions, Position{move, 0, move, direction})
		}
	}
	return positions
}

func moveSubmarine(movements []Position) Position {
	currentPosition := Position{0, 0, 0, ""}
	for _, move := range movements {
		currentPosition.add(move)
	}
	return currentPosition
}

func moveSubmarine2(movements []Position) Position {
	currentPosition := Position{0, 0, 0, ""}
	for _, move := range movements {
		currentPosition.add2(move)
	}
	return currentPosition

}
