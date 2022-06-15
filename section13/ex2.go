package main

import "fmt"
type person struct{
	first string
	last string
	number []int
}

func main(){
	

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
	
	x := map[string]person{
		p1.last : p1,
		p2.last : p2,
		p2.first : p1,
	}
	
	// fmt.Println(p1)
	// fmt.Println(p2)
	// fmt.Println(p1.first,p1.last,p1.number)
	// fmt.Println(p2.first,p2.last,p2.number)
	fmt.Println(x)

	for i,j := range x {
		for _ , j1 := range j.number {
			fmt.Println(i , j1 )
		}
	}
}