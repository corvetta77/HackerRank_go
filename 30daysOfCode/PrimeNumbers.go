package main

import (
	"fmt"
)

func mainp() {
	readNumbersAndPrintResult()
}

func readNumbersAndPrintResult() error {
	var numberOfLines int
	_, err := fmt.Scanf("%d\n", &numberOfLines)
	if err != nil {
		return err
	}
	var lists = make([]int, numberOfLines)
	for i := 0; i < numberOfLines; i++ {
		var number int
		_, err1 := fmt.Scanf("%d\n", &number)
		if err1 != nil {
			return err1
		}
		lists[i] = number
	}
	for i := 0; i < numberOfLines; i++ {
		fmt.Printf("%s\n", isPrime(lists[i]))
	}
	return nil
}

func isPrime(number int) string {
	if number == 1 {
		return "Not prime"
	}
	if number == 2 {
		return "Prime"
	}
	if number%2 == 0 {
		return "Not prime"
	}
	for i := 3; i < number; i += 2 {
		if number%i == 0 {
			return "Not prime"
		}
	}
	return "Prime"
}
