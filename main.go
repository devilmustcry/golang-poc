package main

import (
	"fmt"
)

type QueueImpl struct {
	data []rune
}

func (s *QueueImpl) Pop() rune {
	var pop rune
	s.data, pop = s.data[1:], s.data[0]
	return pop
}

func (s *QueueImpl) Peek() rune {
	var peek rune
	peek = s.data[0]
	return peek
}

func (s *QueueImpl) Append(a rune) {
	s.data = append(s.data, a)
}

func (s *QueueImpl) Size() int {
	return len(s.data)
}

type StackImpl struct {
	data []rune
}

func (s *StackImpl) Pop() rune {
	var pop rune
	s.data, pop = s.data[0:len(s.data)-1], s.data[len(s.data)-1]
	return pop
}

func (s *StackImpl) Peek() rune {
	var peek rune
	peek = s.data[len(s.data)-1]
	return peek
}

func (s *StackImpl) Append(a rune) {
	s.data = append(s.data, a)
}

func (s *StackImpl) Size() int {
	return len(s.data)
}

func main() {
	fmt.Println(isValid("([])"))
}

func isValid(s string) bool {
	openQueue := QueueImpl{}
	closedStack := StackImpl{}
	for _, data := range s {
		switch data {
		case '(', '[', '{':
			openQueue.Append(data)
		case ')', ']', '}':
			closedStack.Append(data)
		}
	}
	if openQueue.Size() != closedStack.Size() {
		return false
	}

	for openQueue.Size() > 0 {
		openPop := openQueue.Pop()
		closedPop := closedStack.Pop()
		if (openPop == '(' && closedPop == ')') ||
			(openPop == '{' && closedPop == '}') ||
			(openPop == '[' && closedPop == ']') {
			continue
		}
		return false
	}
	return true
}
