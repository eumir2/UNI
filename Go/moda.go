package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var s []int
	for i := 1; i < len(os.Args); i++ {
		value, _ := strconv.Atoi(os.Args[i])
		s = append(s, value)
	}
	fmt.Print(moda(s))
}

func moda(s []int) []int {
	var check bool = true
	var mx int = 0
	var count int
	var max []int
	for i, x := range s {
		count = 0
		for j, y := range s {
			if x == y && i != j {
				count++
			}
		}
		if count == mx {
			for _, z := range max {
				if x == z {
					check = false
				}
			}
			if check {
				mx = count
				max = append(max, x)
			}
		} else if count > mx {
			for _, z := range max {
				if x == z {
					check = false
				}
			}
			if check {
				mx = count
				max = nil
				max = append(max, x)
			}
		}

	}
	return max
}
