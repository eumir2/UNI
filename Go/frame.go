package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	size, _ := strconv.Atoi(os.Args[1])
	x, _ := strconv.Atoi(os.Args[2])
	y, _ := strconv.Atoi(os.Args[3])

	grid(size, x, y)
}
func grid(dim int, x int, y int) {
	size := dim + 1
	grid := make([][]string, size)
	for i := 0; i < size; i++ {
		grid[i] = make([]string, size)
		for j := 0; j < size; j++ {
			fmt.Print(grid[i][j])
			if i == x && j == y {
				grid[i][j] = "*"
			} else {
				grid[i][j] = " "
			}
		}
	}
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if i == 0 {
				grid[i][j] = strconv.Itoa(j)
			}
			if j == 0 {
				grid[i][j] = strconv.Itoa(i)
			}
			fmt.Print(grid[i][j])
		}
		fmt.Println()
	}
}
