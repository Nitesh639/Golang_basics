package main

import "fmt"

func main(){
	fmt.Println(foo()())
}

func foo() func() int{
	return func () int {
		return 34
	}
}