package main

import "fmt"

func main() {
	c := make(chan int)
	c2 := make(chan int)
	go gen()
	
	recive(c,c2)
	// recive()
}
func recive(c,c2 chan int) {
	for v := range c {
		fmt.Print(v)
	}
	fmt.Println("")
	for v2 := range c2 {
		fmt.Print(v2)
	}
	fmt.Println("")
}

func gen(c chan int,c2 chan int){
		for i := 0; i < 20; i++ {
		if i%2 ==0{c <- i
		}else{
		c2 <- i
		}
		}
		close(c)
}
