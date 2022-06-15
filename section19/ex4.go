package main

import (
	"fmt"
	"sort"
)

func main(){
	x := []int{23,1,3,54,32,7,54,6,44,21,3,4,6,2,0}
	s := []string{"a","s","d","f","g","h"}

	sort.Ints(x)
	sort.Strings(s)

	fmt.Println(s)
	fmt.Println(x)
}