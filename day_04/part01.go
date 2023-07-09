package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isFullyContained(section1, section2 string) bool {
	chapters1, chapters2 := strings.Split(section1, "-"), strings.Split(section2, "-")

	if (toInt(chapters1[0]) >= toInt(chapters2[0]) && toInt(chapters1[1]) <= toInt(chapters2[1])) ||
		(toInt(chapters2[0]) >= toInt(chapters1[0]) && toInt(chapters2[1]) <= toInt(chapters1[1])) {
		return true
	}
	return false
}

func main() {
	input, _ := os.Open("./input.txt")

	scanner := bufio.NewScanner(input)

	fully_contained_sections := 0

	for scanner.Scan() {
		sections := strings.Split(scanner.Text(), ",")
		if isFullyContained(sections[0], sections[1]) {
			fully_contained_sections += 1
		}
	}

	fmt.Println(fully_contained_sections)
}

func toInt(arg interface{}) int {
	var val int
	switch arg.(type) {
	case string:
		var err error
		val, err = strconv.Atoi(arg.(string))
		if err != nil {
			panic("error converting string to int " + err.Error())
		}
	default:
		panic(fmt.Sprintf("unhandled type for int casting %T", arg))
	}
	return val
}
