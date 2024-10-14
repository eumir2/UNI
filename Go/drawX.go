package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	altezza, _ := strconv.Atoi(os.Args[1])
	StampaX(altezza)
}

func StampaX(alt int) {
	var midspace int = alt - 2
	for startspace := 0; startspace <= (alt / 2); startspace++ {
		if midspace > 0 {
			Dx(startspace)
			Sx(midspace)
		} else if midspace == -1 {
			for i := 0; i <= (alt/2)-1; i++ {
				fmt.Print(" ")
			}
			fmt.Print("*")
		}
		midspace -= 2
		fmt.Println()
	}
	midspace = 1
	for startspace := (alt / 2) - 1; startspace >= 0; startspace-- {
		Dx(startspace)
		Sx(midspace)
		midspace += 2
		fmt.Println()
	}
}

func Dx(space int) {
	for i := 0; i < space; i++ {
		fmt.Print(" ")
	}
	fmt.Print("*")
}

func Sx(space int) {
	for i := space; i > 0; i-- {
		fmt.Print(" ")
	}
	fmt.Print("*")
}
