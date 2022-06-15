package main

import "fmt"

func main(){
	x := struct{
		name string
		number int
	}{
		name : "Nitesh",
		number : 23,
	} 
	fmt.Println(x)
}