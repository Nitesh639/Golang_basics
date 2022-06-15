package main

import "fmt"

type squre struct{
	length float64
}
type circle struct{
	radius float64
}

func (c circle) area() float64{
	return (3.14*c.radius*c.radius)
}

func (s squre) area() float64{
	return (s.length*s.length)
}

type shape interface{
		area() float64
}

func info (s shape){
	fmt.Println(s.area())
}

func main(){
	c1 := circle{
		radius: 4.8,
	}
	s1 := squre{
		length: 2.1,
	}
	fmt.Println(c1.area())
	fmt.Println(s1.area())
	info(c1)
	info(s1)
}