package main

import (
	"fmt"
	"sync/atomic"
	// "runtime"
	"sync"
)

func main(){
	gs := 100
	var  wg sync.WaitGroup
	var incrment int64
	wg.Add(gs)
	for i:=0; i<gs;i++{
		go func ()  {
			atomic.AddInt64(&incrment,1)
			// runtime.Gosched()
			fmt.Println("Incrments\t",atomic.LoadInt64(&incrment))
			wg.Done()
		}()
	}
	// fmt.Println("2nd last",incrment)
	wg.Wait()
	fmt.Println("Last one",incrment)
}