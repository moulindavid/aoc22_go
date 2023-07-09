package main

import (
	"bufio"
	"fmt"
	"os"
)

type stack struct {
	elements []rune
}

func (s *stack) push(r rune) {
	s.elements = append(s.elements, r)
}

func (s *stack) pop() rune {
	r := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return r
}

func (s *stack) addToBottom(r rune) {
	s.elements = append([]rune{r}, s.elements...)
}

func (s stack) String() string {
	var str string
	for _, r := range s.elements {
		str += string(r) + " "
	}
	return str
}

func main() {
	input, _ := os.Open("./input.txt")

	scanner := bufio.NewScanner(input)

	stacks := make([]stack, 9)

	//Scan util we found the crate numbers
	for scanner.Scan() && scanner.Text() != " 1   2   3   4   5   6   7   8   9 " {
		for i, r := range scanner.Text() {
			if r != ' ' && r != '[' && r != ']' {
				//add to the correct stack ( 3 char + 1 space)
				stacks[i/4].addToBottom(r)
			}
		}
	}
	scanner.Scan()

	for scanner.Scan() {
		var toMove, from, to int
		fmt.Sscanf(scanner.Text(), "move %d from %d to %d", &toMove, &from, &to)

		temp_stack := stack{
			elements: []rune{},
		}
		for move := 0; move < toMove; move++ {
			temp_stack.push(stacks[from-1].pop())
		}
		//move elements from one stack to another
		for move := 0; move < toMove; move++ {
			stacks[to-1].push(temp_stack.pop())
		}
	}
	for _, s := range stacks {
		fmt.Print(string(s.pop()))
	}
}
