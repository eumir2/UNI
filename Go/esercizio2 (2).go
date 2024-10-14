package main

import "fmt"

func main() {
	var n, L, H int
	fmt.Scan(&n, &L, &H)
	PrintSnake(n, L, H)
}

func PrintSnake(n int, l int, h int) {
	var value int
	var reverse bool = true
	var reverse2 bool = true
	//nrow := (n * h) - (n - 1)
	for rows := 0; rows < n; rows++ {
		if reverse {
			value, reverse = Row(value, l, reverse)
		} else {
			value, reverse = ReverseRow(value, l, reverse)
		}
		if reverse2 {
			value, reverse2 = Col(value, h-2, l, reverse2)
		} else {
			value, reverse2 = ReverseCol(value, h-2, l, reverse2)
		}
	}
	value, reverse = Row(value, l, reverse)
}

func Row(n int, l int, reverse bool) (int, bool) {
	for i := 0; i < l; i++ {
		n = Count(n)
		fmt.Print(n)
		n++
	}
	reverse = false
	return n, reverse
}

func Col(n int, h int, l int, reverse bool) (int, bool) {
	for i := 0; i < h; i++ {
		fmt.Println()
		_ = Space(l - 1)
		n = Count(n)
		fmt.Print(n)
		n++
	}
	fmt.Println()
	reverse = false
	return n, reverse
}

func Count(n int) int {
	if n == 10 {
		n = 0
	}
	return n
}

func Space(n int) int {
	for i := 0; i < n; i++ {
		fmt.Print(" ")
	}
	return n
}

func ReverseRow(value int, l int, reverse bool) (int, bool) {
	n := value + l - 1
	for i := 0; i < l; i++ {
		n = Count(n)
		fmt.Print(n)
		n--
	}
	reverse = true
	return n + 1 + l, reverse
}

func ReverseCol(n int, h int, l int, reverse bool) (int, bool) {
	for i := 0; i < h; i++ {
		fmt.Println()
		n = Count(n)
		fmt.Print(n)
		n++
	}
	fmt.Println()
	reverse = true
	return n, reverse
}
