package main

import "fmt"

func main(){
	defer foo()
	defer bar()
	ton()
}

func foo(){
	fmt.Println("foo foo foo")
}

func bar(){
	fmt.Println("bar bar bar")
}

func ton(){
	fmt.Println("ton ton ton")
}