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
		data := `[{"First":"Niteh","Last":"Kushwaha"},{"First":"Kamal","Last":"Yadav"},{"First":"Amit","Last":"Kumar"},{"First":"Sanu","Last":"Kumar"}]`
		bs := []byte(data)
		fmt.Println(bs)
		User := []Name{}
		err := json.Unmarshal(bs, &User)
		if err != nil{
			fmt.Println("Error")
		}
		fmt.Println(User)
}