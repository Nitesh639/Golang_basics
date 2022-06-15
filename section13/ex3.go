package main

import (
	"fmt"
)

type vehicle struct{
	door int
	col string
}
type truck struct{
	vehicle
	fourwheel bool
}
type sedan struct{
	vehicle
	luxury bool
}

func main(){
	truck1 := truck{
		vehicle: vehicle{ 
			door : 4,
			col : "blue"},
		fourwheel: true,
		}
	sedan1 := sedan{
		vehicle: vehicle{
			door : 8,
			col : "black",
		},
		luxury: false,
	}

	fmt.Println(truck1)
	fmt.Println(sedan1)
}