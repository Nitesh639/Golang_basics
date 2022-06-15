package main

import "fmt"

func main(){
	birth := 2000
	for {
		fmt.Println(birth)
		birth++
		if birth>2022 {
			break
		}
	}
}