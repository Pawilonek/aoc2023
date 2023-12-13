// https://adventofcode.com/2023/day/3

package main

import (
	"bufio"
	"fmt"
	"os"
)

type engineSchematic [][]byte

type engineBuilder struct {
	schematic engineSchematic
}

const (
	COLOR_DEFAULT     = "\033[0m"
	COLOR_GREY        = "\033[30m"
	COLOR_LIGHT_CYAN  = "\033[96m"
	COLOR_LIGHT_GREEN = "\033[92m"
	COLOR_LIGHT_RED   = "\033[91m"
)

func (e *engineBuilder) isNumberNextToSymbol(row int, col int, length int) bool {
	boxRowStart := row - 1
	if boxRowStart < 0 {
		boxRowStart = 0
	}

	boxRowEnd := row + 1
	maxRows := len(e.schematic) - 1 // the max cound be moved to init
	if boxRowEnd > maxRows {
		boxRowEnd = maxRows
	}

	boxColStart := col - 1
	if boxColStart < 0 {
		boxColStart = 0
	}

	boxColEnd := col + length + 1
	maxCols := len(e.schematic[0]) - 1 // the schematc is always aquare
	if boxColEnd > maxCols {
		boxColEnd = maxCols
	}

	for i := boxRowStart; i <= boxRowEnd; i++ {
		for j := boxColStart; j <= boxColEnd; j++ {
			if e.isSymbol(i, j) {
				return true
			}
		}
	}

	return false
}

func (e *engineBuilder) isNumber(row int, col int) bool {
	return e.schematic[row][col] >= '0' && e.schematic[row][col] <= '9'
}

func (e *engineBuilder) isBlank(row int, col int) bool {
	return e.schematic[row][col] == '.'
}

func (e *engineBuilder) isSymbol(row int, col int) bool {
	return !e.isBlank(row, col) && !e.isNumber(row, col)
}

func (e *engineBuilder) calc() {
	sum := 0 

	for row, line := range e.schematic {
		number := 0
		numbers := []int{}
		positionOfTheNumber := -1

		for col, character := range line {
			if e.isNumber(row, col) {
				if positionOfTheNumber == -1 {
					positionOfTheNumber = col
				}

				digit := character - '0'
				number *= 10
				number += int(digit)

				continue
			}

			if positionOfTheNumber > -1 {
				// Just got a full number

				color := COLOR_LIGHT_RED
				if e.isNumberNextToSymbol(row, positionOfTheNumber, col-positionOfTheNumber-1) {
					color = COLOR_LIGHT_GREEN
					numbers = append(numbers, number)
					sum += number
				}

				fmt.Printf("%s%d%s", color, number, COLOR_DEFAULT)

				positionOfTheNumber = -1
				number = 0
			}

			if e.isBlank(row, col) {
				fmt.Print(COLOR_GREY + "." + COLOR_DEFAULT)

				continue
			}

			fmt.Print(COLOR_LIGHT_CYAN + string(character) + COLOR_DEFAULT)
		}

		if positionOfTheNumber > -1 {
			color := COLOR_LIGHT_RED
			if e.isNumberNextToSymbol(row, positionOfTheNumber, len(line)-positionOfTheNumber-1) {
				color = COLOR_LIGHT_GREEN
				numbers = append(numbers, number)
				sum += number
			}

			fmt.Printf("%s%d%s", color, number, COLOR_DEFAULT)
		}

		fmt.Print("  ", numbers)
		fmt.Println()
	}

	fmt.Println("Result:", sum)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)

		return
	}

	defer file.Close()

	schematic := engineSchematic{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		schematic = append(schematic, []byte(line))
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println(err)

		return
	}

	builder := engineBuilder{
		schematic: schematic,
	}

	builder.calc()
}
