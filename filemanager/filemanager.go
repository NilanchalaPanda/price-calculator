package filemanager

import (
	"bufio"
	"fmt"
	"os"
)

func ReadLines() {
	file, err := os.Open("prices.txt")

	if err != nil {
		fmt.Println("COULD NOT OPEN FILE")
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(file)

	// THIS WILL HOLD ALL THE DATA FROM THE FILES ------
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		fmt.Println("Reading the file content failed")
		fmt.Println(err)
		file.Close()
		return
	}
}
