package main

import (
	"fmt"
	"runtime"
	"sync"
)
var wg sync.WaitGroup
func main(){
	fmt.Println("go routines1:",runtime.NumGoroutine())
	fmt.Println("Number of CPU",runtime.NumCPU())
	wg.Add(2)
	go foo()
	go flow()
	doo()
	fmt.Println("go routines 2:",runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("go routines 3:",runtime.NumGoroutine())
}
func foo(){
	for i:=1;i<5;i++{
		fmt.Println("foo")
	}
	wg.Done()
}
func flow(){
	for i:=1;i<5;i++{
		fmt.Println("flow")
	}
	wg.Done()
}
func doo(){
	for i:=1;i<5;i++{
		fmt.Println("doo")
	}
}