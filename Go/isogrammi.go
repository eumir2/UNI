package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner:=bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		line:=scanner.Text()
		if line==""{
			break
		}
		if isIsogram(line){
			fmt.Println(line+": SI")
		}else{
			fmt.Println(line+": NO")
		}
	}
}

func isIsogram(s string)bool{
	for i,char:=range s{
		for j,char2:=range s{
			if char==char2 && i!=j && string(char)!=" " && string(char)!="-"{
				return false
			}
		}
	}
	return true
}

