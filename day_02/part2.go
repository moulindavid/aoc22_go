package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	figuresCodes = map[string]string{
		"A": "R",
		"B": "P",
		"C": "S",
	}
	figureScore = map[string]int{
		"R": 1,
		"P": 2,
		"S": 3,
	}
	expectedResult = map[string]string{
		"X": "L",
		"Y": "D",
		"Z": "w",
	}
	figuresLoseTo = map[string]string{
		"S": "R",
		"R": "P",
		"P": "S",
	}
	figuresWinTo = map[string]string{
		"P": "R",
		"S": "P",
		"R": "S",
	}
)

func calculateScore(a, b string) int {
	p1, result := figuresCodes[a], expectedResult[b]

	if result == "D" {
		return 3 + figureScore[p1]
	} else if result == "L" {
		return figureScore[figuresWinTo[p1]]
	} else {
		return figureScore[figuresLoseTo[p1]] + 6
	}
}

func main() {
	input, _ := os.Open("./input.txt")

	scanner := bufio.NewScanner(input)

	totalScore := 0

	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " ")
		totalScore += calculateScore(tokens[0], tokens[1])
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(totalScore)
}
