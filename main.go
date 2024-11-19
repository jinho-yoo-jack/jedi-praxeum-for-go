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

	var a int
	var p *int
	p = &a
	*p = 29
	fmt.Println("Memory Address ::: ", *p)
	fmt.Println("a is ", *p)
	fmt.Println("a is ", a)

	var a1 int = 10
	//var b1 float32 = 1.1
	p1 := &a1
	var p2 *int = &a1
	fmt.Println("Memory Address ::: ", p1)
	fmt.Printf("p1 == p2 : %v\n", p1 == p2)

	var x int = 3
	var y int = 2
	var px *int = &x
	var py *int = &y
	fmt.Printf("Memory Address ::: %d", *px+*py)
	//fmt.Printf("Memory Address ::: %s, %s", p1, p2)

}
