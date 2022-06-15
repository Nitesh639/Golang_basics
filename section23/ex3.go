package main

import "fmt"

func main() {
	c := gen()
	recive(c)
	// recive()
}
func recive(c <-chan int) {
	for v := range c {
		fmt.Print(v)
	}
	fmt.Println("")
}

func gen() <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 20; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}
