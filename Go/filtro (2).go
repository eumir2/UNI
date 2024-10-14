package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var rep []int
	var runes []string
	for i := 1; i < len(os.Args); i++ {
		if i%2 == 1 {
			runes = append(runes, os.Args[i])
		} else {
			value, _ := strconv.Atoi(os.Args[i])
			rep = append(rep, value)
		}
	}
	CompressedEncoding(runes, rep)
}
func CompressedEncoding(c []string, s []int) {
	for i, char := range c {
		for j := 0; j < s[i]; j++ {
			fmt.Print(char)
		}
	}
}
