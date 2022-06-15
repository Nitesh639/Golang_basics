package main

import (
	"fmt"
	// "runtime"
	"sync"
)

func main(){
	gs := 100
	var  wg sync.WaitGroup
	incrment := 0
	wg.Add(gs)
	for i:=0; i<gs;i++{
		go func ()  {
			v := incrment
			// runtime.Gosched()
			v++
			incrment = v
			fmt.Println(incrment)
			wg.Done()
		}()
	}
	fmt.Println("2nd last",incrment)
	wg.Wait()
	fmt.Println("Last one",incrment)
}