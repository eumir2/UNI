package main


import "fmt"

func codificaRicorsiva(n int, s string)string{
  if n == 0{
    return s
  }
  s += string(n%2)
  //fmt.Println(n%2)
  //fmt.Println(n)
  return codificaRicorsiva(n/2,s)

}


func main(){
  var n int
  fmt.Scan(&n)
  s := codificaRicorsiva(n,"")
  //stampo la stringa al contrario
  for i := len(s)-1; i >= 0; i--{
    fmt.Print(s[i])
  }
  fmt.Println()
}
