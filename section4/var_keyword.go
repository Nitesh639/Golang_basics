package main

import "fmt"
var y = 78
var z int
func main(){
	x := 43
	fmt.Println(x)
	foo()
	fmt.Println(y)
	fmt.Println(z)
}

func foo(){
	fmt.Println(y)
}