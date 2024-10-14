package main

import (
	"fmt"
	"strings"
)

func main() {

	const K = 5
	var parola string
  var primaVocale rune

	for i:=0; i<K; i++ {
		fmt.Scan(&parola)
   		for _, r := range parola {
      		if strings.ContainsRune("aeiou", r){
            primaVocale = r
            fmt.Println(string(primaVocale))
            break
          }else {
            primaVocale = -1
          }
   		}
      if primaVocale == -1 {
        fmt.Println("la parola", parola, "non contiene vocali")
      }
	}
}
