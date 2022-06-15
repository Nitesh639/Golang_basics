package main

import "fmt"
var x int = 43
var y string = "Hello man"
var z string = `Hello girl,
how are you?
are u waiting for me?`
func main(){
	fmt.Println(x)
	fmt.Printf("%T\n",x)
	fmt.Println(y)
	fmt.Printf("%T\n",y)
	fmt.Println(z)
	fmt.Printf("%T\n",z)
}