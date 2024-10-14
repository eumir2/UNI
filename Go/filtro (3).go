package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"unicode"
)

func main() {
	var line, lines string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line = scanner.Text()
		if line == "" {
			break
		}
		lines += line
	}
	s := Conv(lines)
	s2 := make([]int, len(s))
	copy(s2, s)
	sort.Ints(s2)
	s2 = Reverse(s2)
	if Equal(s, s2) {
		fmt.Println("Sequenza nascosta ordinata")
	} else {
		fmt.Println("Sequenza nascosta non ordinata")
	}
}

func Conv(s string) []int {
	var s2 []int
	for _, x := range s {
		if unicode.IsNumber(x) {
			s2 = append(s2, int(x))
		}
	}
	return s2
}

func Equal(s, s2 []int) bool {
	var ok bool
	for i, x := range s {
		if s2[i] != x {
			ok = false
			return false
		} else {
			ok = true
		}
	}
	return ok
}

func Reverse(s []int) []int {
	var s2 []int
	for i := len(s) - 1; i > -1; i-- {
		s2 = append(s2, s[i])
	}
	return s2
}
