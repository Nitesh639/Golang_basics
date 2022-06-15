package main

import "fmt"

func main(){
	f := func ()  {
		fmt.Println("hello from ffffff")
	}
	f()
	g := func (x int)  {
		fmt.Println("Hy you are now", x)
	}
	g(22)
}