package main
import ("fmt"
        "strconv"
        "strings")

func giorniInMese(s string) (int,int) {
  giorni := 0
  p1 := strings.Index(s,"-") + 1

  s = s[p1 : ]
  //fmt.Println(s)

  p2 := strings.Index(s,"-")

  mese, _ := strconv.Atoi(s[ : p2])

  switch mese{
    case 11,4,6,9 :
      giorni = 30
    case 2 :
      giorni = 20
    default :
      giorni = 31
  }

  return giorni,mese
}


func main() {
  var s string

  fmt.Scan(&s)


  g,m := giorniInMese(s)

  fmt.Printf("il mese %d ha %d giorni",m,g)



}
