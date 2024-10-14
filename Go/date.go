package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var lines []string
	lines = leggiT()
	for _, x := range lines {
		_ = Formatta(x)
	}
}
func leggiT() []string {
	var line string
	var lines []string
	f, _ := os.Open("input.txt")
	scan := bufio.NewScanner(f)
	for scan.Scan() {
		line = scan.Text()
		if line == "" {
			break
		}
		lines = append(lines, line)
	}
	return lines
}

func Formatta(s string) string {
	s = strings.Replace(s, `/`, `-`, 3)
	char := []rune(s)
	if len(char) == 8 {
		fmt.Printf("%c%c%c%c/0%c/0%c\n", char[0], char[1], char[2], char[3], char[5], char[7])
	}
	if len(char) == 9 && char[6] == 45 {
		fmt.Printf("%c%c%c%c/0%c/%c%c\n", char[0], char[1], char[2], char[3], char[5], char[7], char[8])
	}
	if len(char) == 9 && char[6] != 45 {
		fmt.Printf("%c%c%c%c/%c%c/0%c\n", char[0], char[1], char[2], char[3], char[5], char[6], char[8])
	}
	if len(char) == 10 {
		fmt.Printf("%c%c%c%c/%c%c/%c%c\n", char[0], char[1], char[2], char[3], char[5], char[6], char[8], char[9])
	}
	return s

}
