package main

import "fmt"

func main() {
	var a int
	var b int

	n, err := fmt.Scanln(a, b)
	if err != nil {
		print(n, a, b)
	}

}
