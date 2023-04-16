package main

import (
	"fmt"
)

func main121() {
	n := 439
	binaryRepresentation := convertDec2Bin(n)
	fmt.Println(binaryRepresentation)
	numberOfOnes := findConsecutiveOnes(binaryRepresentation)
	fmt.Println(numberOfOnes)
}

func findConsecutiveOnes(representation []int) int {
	consecutiveOnesCounter := 0
	maxConsecutiveOnes := 0
	for i := 0; i < len(representation); i++ {
		if representation[i] == 1 {
			consecutiveOnesCounter++
			if consecutiveOnesCounter > maxConsecutiveOnes {
				maxConsecutiveOnes = consecutiveOnesCounter
			}
		} else {
			consecutiveOnesCounter = 0
		}
	}
	return maxConsecutiveOnes
}

func convertDec2Bin(n int) []int {
	var binary []int

	if n == 0 {
		binary = append(binary, 0)
		return binary
	}
	for n > 0 {
		binary = append(binary, n%2)
		n = n / 2
	}
	var reversedBinary []int
	for i := len(binary) - 1; i >= 0; i-- {
		reversedBinary = append(reversedBinary, binary[i])
	}
	return reversedBinary
}
