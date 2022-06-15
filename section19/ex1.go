package main

import (
	"encoding/json"
	"fmt"
)

type Name struct{
	First string
	Last string
}

func main(){
	p1 := Name{"Niteh", "Kushwaha"}
	p2 := Name{"Kamal", "Yadav"}
	p3 := Name{"Amit", "Kumar"}
	p4 := Name{"Sanu", "Kumar"}

	User := []Name{p1,p2,p3,p4}
	fmt.Println(User)

	bs , err := json.Marshal(User)

	fmt.Println(string(bs))
	fmt.Println(err)
}