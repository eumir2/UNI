package main

import "fmt"

/*Si vuole scrivere un programma max_somma_cifre.go che
 legge da standard input una serie di numeri >= 0, terminata da 999,
 trova il numero (escludendo 999)  la cui somma delle cifre è la maggiore
  e stampa tale somma.*/

func main()  {
  var n int
	maxSum := 0
	for n != 999 {
    sTmp := 0
		for n > 0 {
			sTmp += n % 10
      n /= 10
		}
		if sTmp > maxSum {
			maxSum = sTmp
		}
		fmt.Scan(&n)
	}
	fmt.Println(maxSum)
}
