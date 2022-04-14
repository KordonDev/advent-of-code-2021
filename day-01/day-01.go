package main

import (
	"fmt"
	"math"
)

func main() {
	test1 := stringsToInt(readFile("./test-input.txt"))
	input1 := stringsToInt(readFile("./input.txt"))

	fmt.Println("Solution Test 1", countIncreases(test1))
	fmt.Println("Solution 1", countIncreases(input1))

	testSlidingWindowDepths := createSlidingWindow(test1)
	slidingWindowDepths := createSlidingWindow(input1)
	fmt.Println("Solution Test 2", countIncreases(testSlidingWindowDepths))
	fmt.Println("Solution 2", countIncreases(slidingWindowDepths))
}

func countIncreases(depths []int) int {
	depthIncreased := 0
	var lastDepth = math.MaxInt
	for _, depth := range depths {
		if depth > lastDepth {
			depthIncreased += 1
		}
		lastDepth = depth
	}
	return depthIncreased
}

func createSlidingWindow(depths []int) []int {
	var slidingWindowDepths []int
	for i, _ := range depths {
		if i >= 2 {
			slidingWindowDepths = append(slidingWindowDepths, depths[i]+depths[i-1]+depths[i-2])
		}
	}
	return slidingWindowDepths
}
