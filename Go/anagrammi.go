package main
import "fmt"

/*Scrivere un programma anagrammi.go che legge due stringhe
da linea di comando e valuta se le due stringhe sono
una l'anagramma dell'altra (la seconda stringa è formata
da una permutazione dei caratteri della prima)

In particolare il programma è dotato di:

- una funzione
	func isAnagram(p1, p2 string) bool
che restituisce true se le due stringhe sono una l'anagramma dell'altra,
false altrimenti*/

func isAnagram(p1,p2 string)bool  {
  var slice1,slice2 []rune
  slice1 = make([]rune, len(p1))
  slice2  = make([]rune, len(p2))

  //metto le stringhe nelle loro slice
  for i := 0; i < len(p1)-1; i++{
    slice1[i] = rune(p1[i])
    slice2[i] = rune(p2[len(p2)-1-i])
    }

  //confronto le slice
  for i := range p1{
    if slice1[i] != slice2[i]{
      return false
    }
  }
  return true
}

func main()  {
  var p1,p2 string

  fmt.Scan(&p1, &p2)

  if len(p2) != len(p1){
    fmt.Println("input errato")
  }else{
    if isAnagram(p1,p2){
      fmt.Println("p1 e p2 sono anagrammi")
    }else{
      fmt.Println("p1 e p2 non sono anagrammi")
    }
  }


}
