package main

import (
	"fmt"
	"strings"
)

func main() {
	test := readFile("./test-input.txt")

	fmt.Println("Solution Test 1", fishFarming(test[0], 80))

	input := readFile("./input.txt")
	fmt.Println("Solution 1", fishFarming(input[0], 80))

	fmt.Println("Solution Test 2", fishFarming(test[0], 256))
	fmt.Println("Solution 2", fishFarming(input[0], 256))
}

func fishFarming(start string, daysOfFarming int) int {
	numberStrings := strings.Split(start, ",")
	currentFish := 0
	currentDay := make(map[int]int)
	for _, numberStrings := range numberStrings {
		number := stringToInt(numberStrings)
		count, exists := currentDay[number]
		if exists {
			currentDay[number] = count + 1
		} else {
			currentDay[number] = 1
		}
		currentFish = currentFish + 1
	}

	for i := 0; i < daysOfFarming; i++ {
		nextDay := make(map[int]int)
		for i := 0; i <= 8; i++ {
			count, exists := currentDay[i+1]
			if exists {
				nextDay[i] = count
			} else {
				nextDay[i] = 0
			}
		}
		nextDay[6] = nextDay[6] + currentDay[0]
		nextDay[8] = currentDay[0]
		currentFish = currentFish + currentDay[0]
		currentDay = nextDay
	}
	return currentFish
}
