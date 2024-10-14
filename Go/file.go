package main

import (
	. "fmt"
	"os"
	"strconv"
)

func EPrimo(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	inp := os.Args[1]

	var numPrimi []int

	for i := 0; i < len(inp)-2; i++ {
		Println("tolgo 3", inp[i:len(inp)-len(inp)+i+3])
		num, _ := strconv.Atoi(Sprint(inp[0:i] + inp[len(inp)-len(inp)+i+3:]))
		Println("ottengo: ", num, EPrimo(num))
		if EPrimo(num) {
			numPrimi = append(numPrimi, num)
		}
	}

	for i := 0; i < len(inp)-1; i++ {
		Println("tolgo 2", inp[i:len(inp)-len(inp)+i+2])
		num, _ := strconv.Atoi(Sprint(inp[0:i] + inp[len(inp)-len(inp)+i+2:]))
		Println("ottengo: ", num, EPrimo(num))
		if EPrimo(num) {
			numPrimi = append(numPrimi, num)
		}
	}

	for i := 0; i < len(inp); i++ {
		Println("tolgo 1", inp[i:len(inp)-len(inp)+i+1])
		num, _ := strconv.Atoi(Sprint(inp[0:i] + inp[len(inp)-len(inp)+i+1:]))
		Println("ottengo: ", num, EPrimo(num))
		if EPrimo(num) {
			numPrimi = append(numPrimi, num)
		}
	}
	Println(numPrimi)
}
