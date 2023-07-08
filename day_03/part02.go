package main

import (
	"bufio"
	"fmt"
	"os"
)

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
		itemsFirst := createSetOfItems(scanner.Text())
		scanner.Scan()
		itemsSecond := createSetOfItems(scanner.Text())
		scanner.Scan()

		for _, itemOfThirdElf := range scanner.Text() {
			if itemsSecond[byte(itemOfThirdElf)] && itemsFirst[byte(itemOfThirdElf)] {
				prioritiesSum += int(calculateCharPriority(byte(itemOfThirdElf)))
				break
			}
		}
	}

	fmt.Println(prioritiesSum)
}

func createSetOfItems(rucksack string) map[byte]bool {

	set := make(map[byte]bool)

	for i := 0; i < len(rucksack); i++ {
		set[rucksack[i]] = true
	}

	return set
}
