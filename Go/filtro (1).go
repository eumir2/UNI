package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	euro, _ := strconv.Atoi(os.Args[1])
	fmt.Println("Euro: ", euro)
	Tagliomonete(euro)
}

func Tagliomonete(euro int) {
	var count int
	for euro >= 100 {
		count++
		euro -= 100
	}
	fmt.Println("Taglio: 100 Quantità: ", count)
	count = 0
	for euro >= 50 {
		count++
		euro -= 50
	}
	fmt.Println("Taglio: 50 Quantità: ", count)
	count = 0
	for euro >= 20 {
		count++
		euro -= 20
	}
	fmt.Println("Taglio: 20 Quantità: ", count)
	count = 0
	for euro >= 10 {
		count++
		euro -= 10
	}
	fmt.Println("Taglio: 10 Quantità: ", count)
	count = 0
	for euro >= 5 {
		count++
		euro -= 5
	}
	fmt.Println("Taglio: 5 Quantità: ", count)
	count = 0
	for euro >= 2 {
		count++
		euro -= 2
	}
	fmt.Println("Taglio: 2 Quantità: ", count)
	count = 0
	for euro >= 1 {
		count++
		euro -= 1
	}
	fmt.Println("Taglio: 1 Quantità: ", count)
}
