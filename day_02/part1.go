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
		"X": "R",
		"Y": "P",
		"Z": "S",
	}
	figureScore = map[string]int{
		"R": 1,
		"P": 2,
		"S": 3,
	}
)

func calculateScore(a, b string) int {
	p1, p2 := figuresCodes[a], figuresCodes[b]

	roundScore := figureScore[p2]
	//draw
	if p1 == p2 {
		roundScore += 3
	} else if p1 == "R" && p2 == "S" || p1 == "P" && p2 == "R" || p1 == "S" && p2 == "P" {
		// lose
		roundScore += 0
	} else {
		// win
		roundScore += 6
	}

	return roundScore
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
