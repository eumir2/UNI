package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	fmt.Println("TESTO DI PROVA INSERITO SU PIU' RIGHE")
	LeggiTesto()
}

func LeggiTesto() {
	var mappa map[rune]int
	var testo string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		testo += scanner.Text() + "\n"
	}
	mappa = Occorrenze(testo)
	for k, v := range mappa {
		value := print(v)
		fmt.Printf("%c: %s", k, value)
		fmt.Println()
	}
}

func Occorrenze(s string) map[rune]int {
	testoInrune := []rune(s)
	mappa := make(map[rune]int)
	var slicenum []int
	var slicerun []rune
	fmt.Println(s)
	for _, x := range testoInrune {
		if unicode.IsLetter(x) {
			slicerun = append(slicerun, x)
		}
		if unicode.IsDigit(x) {
			y, _ := strconv.Atoi(string(x))
			slicenum = append(slicenum, y)
		}
	}
	for i := 0; i < len(slicenum); i++ {
		mappa[slicerun[i]] = slicenum[i]
	}
	return mappa
}

func print(val int) string {
	var lol string
	for i := 0; i < val; i++ {
		lol += "*"
	}
	return lol
}
