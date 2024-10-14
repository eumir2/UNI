package main


import "fmt"


func maggioriZeri(a,b int)int{
  max := 0
  numero := a
  for i := a; i <= b; i++{
    n := i
    //fmt.Println(n)
    zeri := 0
    for n > 0{
      if n % 10 == 0{
          zeri++
        }
      n /= 10
    }
    if zeri >= max{
      max = zeri
      numero = i
      }
  }
  return numero
}

func main(){
  var a,b int
  fmt.Scan(&a,&b)

  fmt.Println(maggioriZeri(a,b))
}
