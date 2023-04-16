package main

import (
	"bufio"
	"fmt"
	"os"
)

func main2() {
	readSplitAndMap()
}

func readSplitAndMap() error {
	var numberOfLines int
	_, err := fmt.Scanf("%d\n", &numberOfLines)
	if err != nil {
		return err
	}

	dictionary := make(map[string]string)
	for i := 0; i < numberOfLines; i++ {
		var key string
		var value string
		_, err1 := fmt.Scanf("%s %s\n", &key, &value)
		if err1 != nil {
			return err1
		}
		dictionary[key] = value
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		line := scanner.Text()

		if dictionary[line] != "" {
			fmt.Printf("%s=%s\n", line, dictionary[line])
		} else {
			fmt.Println("Not found")
		}
		// break the loop if line is empty
		if len(line) == 0 {
			break
		}
	}
	return nil
}
