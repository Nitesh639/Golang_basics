package main

import "fmt"
var x bool
func main(){
	fmt.Println(x)
	x = true
	fmt.Println(x)
	a := 23
	b := 4
	fmt.Println(a < b)
	fmt.Println(a > b)
	fmt.Println(a == b)
	fmt.Println(a != b)
}