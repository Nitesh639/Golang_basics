package main

import "fmt"
type Nitesh int
var x Nitesh
func main(){
	fmt.Println(x)
	fmt.Printf("%T\n",x)
	x = 42
	fmt.Println(x)
}