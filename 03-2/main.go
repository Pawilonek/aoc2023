// https://adventofcode.com/2023/day/3

package main

import (
	"bufio"
	"fmt"
	"os"
)

type engineSchematic [][]byte
type position struct {
	col int
	row int
}

func (p *position) String() string {
	return fmt.Sprintf("%dx%d", p.row, p.col)
}

const (
	COLOR_DEFAULT     = "\033[0m"
	COLOR_GREY        = "\033[30m"
	COLOR_LIGHT_CYAN  = "\033[96m"
	COLOR_LIGHT_GREEN = "\033[92m"
	COLOR_LIGHT_RED   = "\033[91m"
)

type engineBuilder struct {
	schematic          engineSchematic
	numbersNextToGears map[position][]int
}

func (e *engineBuilder) getAdjacentGears(row int, col int, length int) []position {
	gears := []position{}

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

	boxColEnd := col + length
	maxCols := len(e.schematic[0]) - 1 // the schematc is always aquare
	if boxColEnd > maxCols {
		boxColEnd = maxCols
	}

	for i := boxRowStart; i <= boxRowEnd; i++ {
		for j := boxColStart; j <= boxColEnd; j++ {
			if e.isGear(i, j) {
				p := position{
					row: i,
					col: j,
				}

				gears = append(gears, p)
			}
		}
	}

	return gears
}

func (e *engineBuilder) isNumber(row int, col int) bool {
	return e.schematic[row][col] >= '0' && e.schematic[row][col] <= '9'
}

func (e *engineBuilder) isGear(row int, col int) bool {
	return e.schematic[row][col] == '*'
}

func (e *engineBuilder) isBlank(row int, col int) bool {
	return !e.isNumber(row, col) && !e.isGear(row, col)
}

func (e *engineBuilder) addToGearsMap(number int, row int, col int) {
	length := len(fmt.Sprintf("%d", number))
	gears := e.getAdjacentGears(row, col, length)

	color := COLOR_LIGHT_RED
	if len(gears) > 0 {
		color = COLOR_LIGHT_GREEN
	}

	fmt.Printf("%s%d%s", color, number, COLOR_DEFAULT)

	for _, gear := range gears {
		e.numbersNextToGears[gear] = append(e.numbersNextToGears[gear], number)
	}
}

func (e *engineBuilder) calc() {
	e.numbersNextToGears = map[position][]int{}

	for row, line := range e.schematic {
		number := 0
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
				e.addToGearsMap(number, row, positionOfTheNumber)
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
			e.addToGearsMap(number, row, positionOfTheNumber)
		}

		fmt.Println()
	}

	gearRatio := 0
	for _, numbers := range e.numbersNextToGears {
		if len(numbers) != 2 {
			continue
		}

		gear := numbers[0] * numbers[1]
		gearRatio += gear
	}

	fmt.Println("Gear ratio:", gearRatio)
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
