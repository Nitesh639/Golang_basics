package main

import "fmt"

func main(){
	c := make(chan int)
	go func ()  {
		c <- 42
	}()
	c2 := make(chan int , 1)
	c2 <- 34
	fmt.Println(<- c)
	fmt.Println(<- c2)
}