package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input, _ := os.Open("./input.txt")

	scanner := bufio.NewScanner(input)

	scanner.Scan()

	signal := scanner.Text()
	for i := 3; i < len(signal); i++ {
		lastFourMap := map[string]int{}
		for j := i - 3; j <= i; j++ {
			lastFourMap[string(signal[j])] += 1
		}
		if len(lastFourMap) == 4 {
			fmt.Print(i + 1)
			break
		}
	}
}
