// https://adventofcode.com/2023/day/2

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cubesSet map[string]int

func calculateThePowerOfGames(games []cubesSet) int {
	minimumCubes := cubesSet{}

	for _, game := range games {
		for color, amount := range game {
			minimumForClor, ok := minimumCubes[color]
			if !ok {
				minimumCubes[color] = amount

				continue
			}

			if amount > minimumForClor {
				minimumCubes[color] = amount
			}
		}
	}

	power := 1
	for _, amount := range minimumCubes {
		power *= amount
	}

	return power
}

type GameParset struct{}

func parseLine(line string) (int, []cubesSet, error) {
	splittedLine := strings.Split(line, ": ")
	if len(splittedLine) != 2 {
		return 0, []cubesSet{}, fmt.Errorf("Too maany `:` in the line.")
	}

	var gameId int
	matched, err := fmt.Sscanf(splittedLine[0], "Game %d", &gameId)
	if err != nil {
		return 0, []cubesSet{}, err
	}

	if matched != 1 {
		return 0, []cubesSet{}, fmt.Errorf("Matched %d elements instead 1.", matched)
	}

	games := []cubesSet{}
	gamesText := strings.Split(splittedLine[1], "; ")
	if len(gamesText) < 1 {
		return 0, []cubesSet{}, fmt.Errorf("There are no games in this line.")
	}

	for _, cubesText := range gamesText {
		oneHandCubes := strings.Split(cubesText, ", ")
		handCubesSet := cubesSet{}
		for _, cubes := range oneHandCubes {
			var amount int
			var color string
			fmt.Sscanf(cubes, "%d %s", &amount, &color)
			handCubesSet[color] = amount
		}

		games = append(games, handCubesSet)
	}

	return gameId, games, nil
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)

		return
	}

	defer file.Close()

	sumOfPowers := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		_, games, err := parseLine(line)
		if err != nil {
			fmt.Println(err)
			return
		}

		power := calculateThePowerOfGames(games)
		sumOfPowers += power
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(sumOfPowers)
}
