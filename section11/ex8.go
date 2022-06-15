package main

import "fmt"

func main(){
	x := map[string][]int{"name":[]int{1,2,3,4},"Nites":{1,5,3,4},"hello":{1,2,9,4},"kumar":{1,2,3,0},}
	for i,j := range x{
		fmt.Printf("%T\t%T\n",j,i)
	}
}