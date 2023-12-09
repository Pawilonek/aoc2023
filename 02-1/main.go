// https://adventofcode.com/2023/day/2

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	COLOR_BLUE  = "blue"
	COLOR_GREEN = "green"
	COLOR_RED   = "red"
)

type cubesSet map[string]int

type Game struct {
	startingCubes cubesSet
}

func (g *Game) isItPossibleToPlay(cubes cubesSet) bool {
	for color, amount := range cubes {
		maxAmount, ok := g.startingCubes[color]
		if !ok {
			return false
		}

		if amount > maxAmount {
			return false
		}
	}

	return true
}

func (g *Game) areTheGamesPossibleToPlay(games []cubesSet) bool {
	for _, game := range games {
		possible := g.isItPossibleToPlay(game)
		if !possible {
			return false
		}
	}

	return true
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
	gameEngine := Game{
		startingCubes: cubesSet{
			COLOR_RED:   12,
			COLOR_GREEN: 13,
			COLOR_BLUE:  14,
		},
	}

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)

		return
	}

	defer file.Close()

	sumOfPossibleGames := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		gameId, games, err := parseLine(line)
		if err != nil {
			fmt.Println(err)
			return
		}

		possible := gameEngine.areTheGamesPossibleToPlay(games)
		if !possible {
			continue
		}

		sumOfPossibleGames += gameId
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(sumOfPossibleGames)
}
