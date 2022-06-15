package main

import "fmt"

func main(){
	y := [5]int{23,42,12,23}
	var x [5]int
	fmt.Println(x)
	x[0]=1
	x[1]=3
	x[2]=12
	x[3]=23
	x[4]=10
	fmt.Println(x)
	fmt.Printf("%T\t%v\n",x,x)
	fmt.Printf("%T\t%v",y,y)
}