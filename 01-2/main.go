package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numberNames = map[int]string{
	1: "one",
	2: "two",
	3: "three",
	4: "four",
	5: "five",
	6: "six",
	7: "seven",
	8: "eight",
	9: "nine",
}

func getFirstAndLastNumber(line string) (int, int) {
	first := -1
	last := -1

	// Track how much letterf of a number we already found
	foundNumberNames := map[int]int{}
	for i := 0; i < len(line); i++ {
		number, err := strconv.Atoi(string(line[i]))
		if err == nil {
			if first < 0 {
				first = number
			}

			last = number

			continue
		}

		// Try to match letter to a number
		for intValue, nameOfANumber := range numberNames {
			position, ok := foundNumberNames[intValue]
			if !ok {
				position = 0
			}

			// Run out of letters in a number - should not happen
			if len(nameOfANumber) < position {
				foundNumberNames[intValue] = 0

				continue
			}

			expectedLetter := nameOfANumber[position]

			// Check if we match a letter for current value
			if expectedLetter != line[i] {
				foundNumberNames[intValue] = 0

				continue
			}

			// Check if we fount the number
			if len(nameOfANumber) == position+1 {
				foundNumberNames[intValue] = 0
				if first < 0 {
					first = intValue
				}

				last = intValue

				break
			}

			foundNumberNames[intValue] = position + 1
		}
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
		number := first*10 + last
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
