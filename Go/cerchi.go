package main

import (
	"fmt"
	"bufio"
	"io"
	"os"
	"strings"
)

func main() {
	b:=bufio.NewScanner(os.Stdin)
	word:=strings.Split(b)
	for b.Scan(){
		splittedline:=strings.Fields(b.Text())
		if splittedline==""{
			break
		}

	}
}
