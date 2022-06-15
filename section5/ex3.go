package main

import "fmt"

var x int = 42
var y string ="Hello"
var z bool = true

func main(){
	fmt.Println(x)
	s := fmt.Sprintf("%T\n%T\n%v\n",x,y,z)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(s)
}