package main

import "fmt"

func main(){
	x:=[]string{"Nitesh","Kumar","Kushwaha"}
	y:=[]string{"Hello", "World"}

	z :=[][]string{}

	z = append(z, x,y)
	fmt.Println(z)
	for i1,v := range z{
		for i2,val := range v{
			fmt.Println(i1,i2,val)
		}
	}
}