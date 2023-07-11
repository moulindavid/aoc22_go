package main

import (
	"bufio"
	"fmt"
	"os"
)

func findMarker(signal string, char_num int) int {

	for i := char_num - 1; i < len(signal); i++ {
		lastNMap := map[string]int{}
		for j := i - (char_num - 1); j <= i; j++ {
			lastNMap[string(signal[j])] += 1
		}
		if len(lastNMap) == char_num {
			return i + 1
		}
	}
	//Not possible
	return 0
}
func main() {
	input, _ := os.Open("./input.txt")

	scanner := bufio.NewScanner(input)

	scanner.Scan()
	signal := scanner.Text()
	fmt.Println(findMarker(signal, 4))
	fmt.Println(findMarker(signal, 14))
}
