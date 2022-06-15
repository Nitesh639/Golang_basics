package main

import "fmt"

func main() {
	a := foo()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(foo()())
	fmt.Println(foo()())
	fmt.Println(foo()())
}

func foo() func() int {
	x := 0
	return func() int{
		x++
		return x
	}
}