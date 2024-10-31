package main

import (
	"fmt"
	house2 "main/struct_example"
	direction "main/switch_function"
)

func main() {

	d := direction.GetDirection(135)
	fmt.Println(d)

	var temp [3]int = [3]int{1, 2, 3}
	t := [...]string{"1", "2", "3"}
	for v := range t {
		fmt.Println("value ::: ", v)
	}
	for i, v := range temp {
		fmt.Println(i, v)
	}
	var house house2.Home = house2.Home{Address: "경기도 군포시 당정동", Size: 29, Price: 9.8, Type: "아파트"}
	fmt.Println(house)
}
