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
	maxCalories1 := 0
	maxCalories2 := 0
	maxCalories3 := 0

	currentElfCalories := 0

	for sc.Scan() {
		snack, err := strconv.Atoi(sc.Text())
		currentElfCalories += snack

		//If error is different from nil, we found an empty line.
		if err != nil {
			if currentElfCalories > maxCalories3 {
				maxCalories3 = currentElfCalories
			}
			if maxCalories3 > maxCalories2 {
				maxCalories3, maxCalories2 = maxCalories2, maxCalories3
			}
			if maxCalories2 > maxCalories1 {
				maxCalories2, maxCalories1 = maxCalories1, maxCalories2
			}
			//End of current elf.
			currentElfCalories = 0
		}
	}
	fmt.Println(maxCalories1 + maxCalories2 + maxCalories3)
}
