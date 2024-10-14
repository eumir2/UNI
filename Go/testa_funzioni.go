package main
import "fmt"

//funzione che controlla se la stringa è palindroma
func isPalindrome(s string) bool{
  var j = len(s)-1
  for i := 0; i < len(s)/2; i++{
    if s[i] != s[j]{
      return false
    }
    j--
  }
  return true
}

//funzione che controlla se la stringa ha maiuscole
func  hasUpper(s string) bool{
  for i := 0; i < len(s); i++{
    if s[i] <= 'Z' && s[i] >= 'A'{
      return true
    }
  }
  return false
}

//funzione che conta cifre ,lettere e altri caratteri in una stringa
func count(s string) (int, int, int){
  cc := 0
  cl := 0
  cr := 0
  for i := 0; i < len(s); i++{
    if (s[i] <= 'Z' && s[i] >= 'A') || (s[i] <= 'z' && s[i] >= 'a'){
      cl++
    }else if s[i] <= '9' && s[i] >= '1'{
      cc++
    }else{
      cr++
    }
  }
  return cc,cl,cr
}


func main()  {
  var s string
  fmt.Scan(&s)

  if isPalindrome(s){
    fmt.Println("palindroma")
  }else{
    fmt.Println("non palindroma")
  }

  if hasUpper(s){
    fmt.Println("ha maiuscole")
  }else{
    fmt.Println("non ha maiuscole")
  }

  cc,cl,cr := count(s)

  fmt.Printf("%d,%d,%d \n",cc,cl,cr)




}
