package main

import "fmt"

type person struct{
	name string
}

func (p person) speak(){
	fmt.Println(p.name)
}

type human interface{
	speak()
}
func saysomething(h human){
	h.speak()
}

func main(){
	p1 := person{
		name: "Nitesh",
	}
	saysomething(&p1)
}