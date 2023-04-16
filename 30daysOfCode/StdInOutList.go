package main

import (
	"fmt"
)

func main1() {
	readSplitAndReorder()
}

func readSplitAndReorder() error {
	var numberOfLines int
	_, err := fmt.Scanf("%d\n", &numberOfLines)
	if err != nil {
		return err
	}
	var lists = make([]string, numberOfLines)
	for i := 0; i < numberOfLines; i++ {
		var line string
		_, err1 := fmt.Scanf("%s\n", &line)
		if err1 != nil {
			return err1
		}
		lists[i] = line
	}
	for i := 0; i < numberOfLines; i++ {
		line := lists[i]
		for j := 0; j < len(line); j = j + 2 {
			fmt.Printf("%c", line[j])
		}
		fmt.Print(" ")

		for j := 1; j < len(line); j = j + 2 {
			fmt.Printf("%c", line[j])
		}
		fmt.Print("\n")
	}
	return nil
}
