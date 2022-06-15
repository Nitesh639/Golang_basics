package main

import "fmt"
type person struct{
	first string
	last string
	age int
}
func main(){
	p1 := person{
		first: "Nitesh",
		last: "kushwaha",
		age: 22,
	}
	p1.speak()
}

func (p person) speak(){
	fmt.Println(p.age,p.first,p.last)
}