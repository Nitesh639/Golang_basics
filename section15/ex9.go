package main

import "fmt"
func foo(x ...int) int{
	total :=0
	for _, i := range x{
		total+= i
	}
	return total
}

func main(){
	m := []int{1,2,3,4,5,6,7,9,12}
	n := even(foo ,m...)
	fmt.Println(n)
}



func even(f func(x ...int) int,y ...int) int{
		var z []int
		for _,j := range y {
			if (j%2 == 0) {
				z = append(z,j)
			}
		}
		return f(z...)
}