package main

import "fmt"

func main(){
	cs := make(chan <-int)
	cr := make(<- chan int)
	go func ()  {
		cs <- 42
	}()
	// fmt.Println("It's only send the values",<-cs)  
	fmt.Println("It's only send the values",<-cr)   //it's only send the values in channel
}