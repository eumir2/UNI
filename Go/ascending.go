package main

import "fmt"

func main() {
	const K = 10

	var prev, next int

  fmt.Scan(&next)
  ascending := true

  for i := 1; i < K; i++ {
    prev = next
    fmt.Scan(&next)
    if next <= prev {
      ascending = false
      break
    }
  }

	if !ascending {
		fmt.Print("non ")
	}
	fmt.Println("crescente")
}
