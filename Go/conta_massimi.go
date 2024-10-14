package main

import "fmt"

/*programma conta_massimi.go che legga una sequenza di 10 interi e
stampi il massimo intero letto
 e quante volte il massimo compare nella sequenza. */

 func main(){
  var number int

  fmt.Scan(&number)
  max := number
  count := 1

  for i := 1; i<10; i++ {
    fmt.Scan(&number)
    if number == max {
      count ++
    }else if number > max{
       max = number
       count = 1
    }
  }
  fmt.Println(max)
  fmt.Println(count)
}
