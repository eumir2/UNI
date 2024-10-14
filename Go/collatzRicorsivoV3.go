package main

import "fmt"

func collatz(n int, p *[]int) *[]int{
  if n == 1{
    *p = append(*p,1)
    return p
  }

  *p = append(*p,n)
  if n % 2 == 0{
    return  collatz(n/2,p)
  }else{
    return  collatz(n * 3 + 1,p)
  }
}

func main()  {
  var n int
  var sequenza = []int{}

  fmt.Scan(&n)

  fmt.Println(collatz(n,&sequenza))
}
