package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input, _ := os.Open("input.txt")
	sc := bufio.NewScanner(input)

	//Search for the maximum sum of calories
	maxCalories := 0
	currentElfCalories := 0

	for sc.Scan() {
		snack, err := strconv.Atoi(sc.Text())
		currentElfCalories += snack

		//If error is different from nil, we found an empty line
		if err != nil {
			if currentElfCalories > maxCalories {
				maxCalories = currentElfCalories
			}
			//End of current elf.
			currentElfCalories = 0
		}
	}
	fmt.Println(maxCalories)
}
