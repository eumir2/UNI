package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var num, molt int
	molt = 1
	for i := 1; i < len(os.Args); i++ {
		num, _ = strconv.Atoi(os.Args[i])
		if num == 0 {
			num = 1
		}
		molt *= num

	}
	fmt.Println(molt)
}
