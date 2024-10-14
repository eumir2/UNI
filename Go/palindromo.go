package main

import "fmt"
import "os"

/*programma che stabilisce se una funzione è palindroma ricorsivamente*/
func isPalindromo(s string) bool{
  if len(s) == 0 || len(s) == 1 {
    return true
  }
  l := len(s)-1
  if s[0] == s[l]{
    return isPalindromo(s[1:l])
  }else{
    return false
  }
}

func main()  {
   stringa := os.Args[1]
   //fmt.Println(stringa)
   if isPalindromo(stringa){
     fmt.Println("s è palindroma")
   }else{
     fmt.Println("s non è palindroma")
   }

}
