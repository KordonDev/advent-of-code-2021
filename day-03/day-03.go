package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	test1 := readFile("./test-input.txt")
	fmt.Println("Solution Test 1", countAndMultiply(test1))

	input1 := readFile("./input.txt")
	fmt.Println("Solution Test 1", countAndMultiply(input1))

	test2 := readFile("./test-input.txt")
	fmt.Println("Solution Test 2", filterAndMultiply(test2))

	input2 := readFile("./input.txt")
	fmt.Println("Solution 2", filterAndMultiply(input2))
}

func mostCommonInColumns(input []string) []string {
	var countOne = make([]int, len(input[0]))
	for columnIndex := range input[0] {
		countOne[columnIndex] = oneInColumn(input, columnIndex)
	}

	var mostCommonNumbers []string
	for _, count := range countOne {
		if count > (len(input) / 2) {
			mostCommonNumbers = append(mostCommonNumbers, "1")
		} else {
			mostCommonNumbers = append(mostCommonNumbers, "0")
		}
	}
	return mostCommonNumbers
}

func oneInColumn(input []string, column int) int {
	var countOne int
	for _, row := range input {
		if row[column] == '1' {
			countOne = countOne + 1
		}
	}
	return countOne
}

func invert(input []string) []string {
	var result []string
	for _, number := range input {
		if number == "0" {
			result = append(result, "1")
		} else {
			result = append(result, "0")
		}
	}
	return result
}

func binaryToInt(input string) int64 {
	i, err := strconv.ParseInt(input, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func countAndMultiply(input []string) int64 {
	gammaList := mostCommonInColumns(input)
	epsilonList := invert(gammaList)
	gamma := binaryToInt(strings.Join(gammaList, ""))
	epsilon := binaryToInt(strings.Join(epsilonList, ""))
	return gamma * epsilon
}

func filterByCharInColumn(input []string, column int, character uint8) []string {
	var result []string
	for _, row := range input {
		if row[column] == character {
			result = append(result, row)
		}
	}
	return result
}

func filterAndMultiply(input []string) int64 {
	columnIndex := 0
	strongOne := make([]string, len(input))
	copy(strongOne, input)
	strongZero := make([]string, len(input))
	copy(strongZero, input)
	for len(strongOne) > 1 {
		if oneInColumn(strongOne, columnIndex)*2 >= len(strongOne) {
			strongOne = filterByCharInColumn(strongOne, columnIndex, '1')
		} else {
			strongOne = filterByCharInColumn(strongOne, columnIndex, '0')
		}
		columnIndex = columnIndex + 1
	}

	columnIndex = 0
	for len(strongZero) > 1 {
		if oneInColumn(strongZero, columnIndex)*2 < len(strongZero) {
			strongZero = filterByCharInColumn(strongZero, columnIndex, '1')
		} else {
			strongZero = filterByCharInColumn(strongZero, columnIndex, '0')
		}
		columnIndex = columnIndex + 1
	}
	strongOneDez := binaryToInt(strongOne[0])
	strongZeroDez := binaryToInt(strongZero[0])
	return strongOneDez * strongZeroDez
}
