package main

import (
	"fmt"
	// "runtime"
	"sync"
)

func main(){
	gs := 100
	var  wg sync.WaitGroup
	var m sync.Mutex
	incrment := 0
	wg.Add(gs)
	for i:=0; i<gs;i++{
		go func ()  {
			m.Lock()
			v := incrment
			v++
			incrment = v
			fmt.Println(incrment)
			m.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Last one",incrment)
}