package main

import (
	"fmt"
	"os"
)

func main() {
	var strings []string
	var s string
	//SIM 1
	/*for i := 1; i < len(os.Args); i++ {
		s += os.Args[i]
	}*/
	//SIM 2
	s = os.Args[1]
	for i := 0; i < len(s); i++ {
		for j := len(s); j > i; j-- {
			ok, word := check(s[i:j], s)
			if ok {
				strings = append(strings, word)
			}
		}
	}
	check2(strings)
}

func check(sl string, s string) (bool, string) {
	if sl[0] == sl[len(sl)-1] && len(sl) > 2 {
		return true, sl
	}
	return false, sl
}

func check2(s []string) {
	var scount []int
	s, scount = removeDuplicateValues(s)
	if len(s) != 0 {
		for i, word := range s {
			fmt.Println(word, " -> Occorrenze: ", scount[i])
		}
	}
}

func removeDuplicateValues(stringSlice []string) ([]string, []int) {
	keys := make(map[string]bool)
	list := []string{}
	var count []int
	var intcount int = 1
	for i, s1 := range stringSlice {
		intcount = 1
		for j, s2 := range stringSlice {
			if s1 == s2 && i != j {
				intcount++
			}
		}
		count = append(count, intcount)
	}
	for _, entry := range stringSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list, count
}
