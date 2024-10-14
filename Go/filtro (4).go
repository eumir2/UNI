package main

import (
	"fmt"
	"os"
	"unicode"
)

func main() {
	var sl []string
	var count int
	for i := 1; i < len(os.Args); i++ {
		word := os.Args[i]
		ok := ParseWord(word)
		if ok {
			sl = append(sl, word)
		} else {
			count++
		}
	}
	Print(sl)
	fmt.Print(" ", count)
}

func ParseWord(s string) bool {
	var ok bool
	for _, char := range s {
		if unicode.IsDigit(char) == false {
			ok = true
		} else {
			return false
		}
	}
	return ok
}

func Print(sl []string) {
	for _, word := range sl {
		fmt.Print(word + " ")
	}
}
