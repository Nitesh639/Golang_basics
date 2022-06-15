package main

import "fmt"

type person struct{
	first string
	last string
}
func main(){
	p1 := person{
		"Nitesh",
		"Kushwaha",
	}

	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}

func changeMe(p *person){
		fmt.Println(p.first)
		p.first = "Shippu"
		fmt.Println(p.first)
}