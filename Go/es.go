package main

import ("fmt"
        "os"
        "strconv")

func f(s[]int, k int)[]int{
  out := []int{}

  occ := map[int]int{}
  for _,v := range s{
    occ[v]++
  }

  for i,v := range occ{
    if v == k{
      out = append(out,i)
    }
  }

  return out
}

func main(){
  k,_ := strconv.Atoi(os.Args[1])
  sliceString := os.Args[2:]

  sliceInt := []int{}
  for _,k := range sliceString{
    val, _ := strconv.Atoi(k)
    sliceInt = append(sliceInt,val)
  }

  fmt.Println(f(sliceInt,k))

}
