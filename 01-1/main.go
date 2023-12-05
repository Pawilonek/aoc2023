package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getFirstAndLastNumber(line string) (int, int) {
	first := -1
	last := -1

	for i := 0; i < len(line); i++ {
		number, err := strconv.Atoi(string(line[i]))
		if err != nil {
			// Error means that this is not a number
			continue
		}

		if first < 0 {
			first = number
		}

		last = number
	}

	return first, last
}

func solve(filename string) (int, error) {
	result := 0

	file, err := os.Open(filename)
	if err != nil {
		return 0, fmt.Errorf("Could not open input file: %w", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		first, last := getFirstAndLastNumber(line)
		number := first * 10 + last
		result += number
	}

	err = scanner.Err()
	if err != nil {
		return 0, fmt.Errorf("Error during reading the file: %w", err)
	}

	return result, nil
}

func main() {
	result, err := solve("input.txt")
	if err != nil {
		fmt.Println("err: %w", err)

		return
	}

	fmt.Println(result)
}
