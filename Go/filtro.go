package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var esp = [11]float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 1; i < len(os.Args); i++ {
		value, _ := strconv.Atoi(os.Args[i])
		for _, x := range esp {
			if math.Pow(3, x) == float64(value) {
				fmt.Print(value, " ")
			}
		}
	}
}
