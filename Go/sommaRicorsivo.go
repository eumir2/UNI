package main

import (
      "fmt"

)

func somma(s []int) int  {
  tmp := 0
  if len(s) <= 0{
    return 0
  }else{
    tmp += s[0]
    fmt.Println(tmp)
    return  tmp + somma(s[1:])
  }
}


func main()  {
  var s = []int {2,3,4,5,6,3,26,2}

  fmt.Println(somma(s))
}
