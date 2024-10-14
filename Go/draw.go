package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func DrawPoint(c byte, l int) string {
	s := strings.Repeat(" ", l) + string(c)
	return s
}

func DrawSegment(c byte, k,l int) string {
	s := strings.Repeat(" ", k) + strings.Repeat(string(c), l)
	return s
}

func PrintScales(l, k int, c byte) {
	for i := 0; i < l; i++ {
		if i == 0 {
			fmt.Println(strings.Repeat(string(c), l))
			for j := 0; j < l-1; j++ {
				fmt.Println(DrawPoint(c, k-1))
			}
		} else {
			fmt.Println( DrawSegment(c,k,l))
			for j := 0; j < l-1; j++ {
				fmt.Println(DrawPoint(c, l))
			}
		}
	}
}

func main() {
	l, _ := strconv.Atoi(os.Args[1])
	n, _ := strconv.Atoi(os.Args[2])
	c := []byte((os.Args[3]))
	s := c[0]
	fmt.Println(DrawSegment('p', 5, 3))
	if l > 0 && n > 0 {
		PrintScales(l, n, s)
	}
}
