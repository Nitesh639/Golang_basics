package main

import "fmt"
const (
	a = 23 + iota	
	b = 23 + iota	
	c = 23 + iota			
	d = 23 + iota	
)
func main(){
	fmt.Println(a,b,c,d)
}