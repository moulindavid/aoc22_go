package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.Open("./input.txt")

	scanner := bufio.NewScanner(input)
	var total_bytes int64
	var stack []int64
	for scanner.Scan() {
		switch line := scanner.Text(); line[:4] {
		case "$ cd":
			if path := line[5:]; path != ".." {
				// we cd in a new dir
				stack = append(stack, 0)
				continue
			}
			// cd out of directory
			dir_size := stack[len(stack)-1]
			if dir_size <= 100_000 {
				total_bytes += dir_size
			}

			if stack = stack[:len(stack)-1]; len(stack) > 0 {
				// directory inside directory
				stack[len(stack)-1] += dir_size
			}
		case "$ ls", "dir ": // hext line
		default:
			// file
			fs := strings.Fields(line)
			fileSize, _ := strconv.ParseInt(fs[0], 10, 64)
			stack[len(stack)-1] += fileSize
		}
	}

	fmt.Println(total_bytes)
}
