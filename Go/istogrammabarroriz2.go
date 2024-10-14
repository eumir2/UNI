package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("TESTO DI PROVA INSERITO SU PIU' RIGHE")
	LeggiTesto()
}

func LeggiTesto() {
	var slicerun []rune
	var slicenum []int
	var mappa map[rune]int
	var testo string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		testo += scanner.Text() + "\n"
	}
	mappa = Occorrenze(testo)
	for k, x := range mappa {
		slicerun = append(slicerun, k)
		slicenum = append(slicenum, x)
	}
	SortRunes(slicerun)
	for i := 0; i < len(slicerun); i++ {
		fmt.Printf("%c: %s", slicerun[i], print(slicenum[i]))
		fmt.Println()
	}
}

func Occorrenze(s string) map[rune]int {
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, "\n", "")
	fmt.Println(s)
	testoInrune := []rune(s)
	mappa := make(map[rune]int)
	var slicerun []rune
	var count int
	var frequency []int
	for _, x := range testoInrune {
		slicerun = append(slicerun, x)
		count = 0
		for _, y := range testoInrune {
			if x == y {
				count++
			}
		}
		frequency = append(frequency, count)
	}
	SortRunes(slicerun)
	for i := 0; i < len(frequency); i++ {
		mappa[slicerun[i]] = frequency[i]
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

func SortRunes(a []rune) {
	for i := 0; i < len(a)-1; i++ {
		indiceMin := i
		for j := i + 1; j < len(a); j++ {
			if a[indiceMin] > a[j] {
				indiceMin = j
			}
		}
		a[i], a[indiceMin] = a[indiceMin], a[i]
	}
}
