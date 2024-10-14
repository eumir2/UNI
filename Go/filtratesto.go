package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	var stringnum []string
	lines := LeggiTesto()
	stringnum = filtratesto(lines)
	for _, x := range stringnum {
		fmt.Println(x, "\n")
	}
}
func LeggiTesto() []string {
	var lines []string
	var stringnum []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	stringnum = filtratesto(lines)
	return stringnum
}
func ContieneCifre(s string) bool {
	for _, x := range s {
		if unicode.IsDigit(x) {
			return true
		}
	}
	return false
}
func filtratesto(lines []string) []string {
	var stringnum []string
	for _, x := range lines {
		if ContieneCifre(x) {
			stringnum = append(stringnum, x)
		}
	}
	return stringnum
}
