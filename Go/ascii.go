package main

import (
	"fmt"
	"unicode"
)

func main() {
	str := "⌘questa ée' una stringa 語formatàa di本 soli ASCII⌘\n"
	fmt.Print(soloASCII(str))
}
func soloASCII(str string) string {
	var newstr string
	for _, x := range str {
		if x > unicode.MaxASCII {
		} else {
			newstr += string(x)
		}
	}
	return newstr
}
