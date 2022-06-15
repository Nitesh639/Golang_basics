package main

import "fmt"

func main(){
	x := 23
	// print the value but not use \n 
	fmt.Print(x)		
	// use only for know the data types and formating the data
	fmt.Printf("%T\n", x)
	// print the value but use \n 
	fmt.Println(x)
	fmt.Printf("%b\n", x)
}