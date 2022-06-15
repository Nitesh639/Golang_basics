package main

import "fmt"

var x int
type hotty int
var y hotty
func main(){
	x = 23
	y = 34
	fmt.Printf("%T\n",x)
	fmt.Printf("%T\n",y)
	fmt.Println(x)
	fmt.Println(y)

	// x = y we can't assing like that
	x = int(y)
	fmt.Printf("%T\n",x)
	fmt.Println(x)

}