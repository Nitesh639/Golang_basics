package main

import "fmt"

func main(){
	type person struct{
		first string
		last string
		number []int
	}

	p1 := person{
		first: "Nitesh",
		last: "Kushwaha",
		number: []int{1,2,3,4},
	}

	p2 := person{
		first: "Hello",
		last: "World",
		number: []int{5,1,3,8},
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p1.first,p1.last,p1.number)
	fmt.Println(p2.first,p2.last,p2.number)
}