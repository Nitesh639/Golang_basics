package main

import "fmt"

func main(){
	ii := []int{1,2,3,4,5,6,7}
	fmt.Println(foo(ii))
	fmt.Println(bar(1,2,3,4,5))
}

func foo(xi []int) int{
	total := 0
	for _, i := range xi {
		total +=i
	}
	return total
}
func bar(xi ...int) int{
	total := 0
	for _, i := range xi {
		total +=i
	}
	return total
}