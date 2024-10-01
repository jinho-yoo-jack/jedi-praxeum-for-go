package main

import "fmt"

func main() {
	// int type
	var x int = 1
	var y int = 2
	c := x + y
	fmt.Println(c)

	// string
	var name string = "jinho"
	fmt.Printf("Owner name is %s", name)

	// boolean
	isTrue := true
	var isFalse = false
	fmt.Println(isTrue, isFalse)
	fmt.Print(isTrue == isFalse)
}
