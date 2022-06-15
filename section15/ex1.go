package main

import "fmt"

func main(){
	fmt.Println(foo())
	fmt.Println(bar())
}

func foo() int{
	return 23
}
func bar() (string,int) {
	return "hello" , 21
}