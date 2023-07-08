package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func findErrors(rucksack string) byte {
	rucksack_len := len(rucksack)
	part1, part2 := rucksack[:rucksack_len/2], rucksack[rucksack_len/2:rucksack_len]

	byteMap := make(map[byte]bool)
	for i := 0; i < len(part1); i++ {
		byteMap[part1[i]] = true
	}

	for i := 0; i < len(part2); i++ {
		if byteMap[part2[i]] {
			return part2[i]
		}
	}
	log.Fatal("we should not be in this case")

	return 'a'
}

// upper char start at 65 so for A to be 27 we need to substract 38
// lower char start at 97 so for a to be 1 we need to substract 96
func calculateCharPriority(error byte) byte {
	if error < 97 {
		return error - 38
	}
	return error - 96
}
func main() {
	input, _ := os.Open("./input.txt")

	scanner := bufio.NewScanner(input)

	prioritiesSum := 0

	for scanner.Scan() {
		rucksacks := scanner.Text()
		curEr := findErrors(rucksacks)
		prioritiesSum += int(calculateCharPriority(curEr))
	}

	fmt.Println(prioritiesSum)
}
