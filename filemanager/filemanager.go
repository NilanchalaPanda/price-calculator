package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("COULD NOT OPEN FILE")
	}

	scanner := bufio.NewScanner(file)

	// THIS WILL HOLD ALL THE DATA FROM THE FILES ------
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		return nil, errors.New("READING THE FILE CONTENT FAILED")
	}

	file.Close()
	return lines, nil
}

func WriteJSON(path string, data interface{}) error {
	file, err := os.Create(path)

	if err != nil {
		return errors.New("FAILED TO WRITE DATA")
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		fmt.Println("FAILED TO CONVERT TO JSON")
	}

	file.Close()
	return nil
}
