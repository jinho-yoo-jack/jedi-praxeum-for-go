package main

import (
	"fmt"
	bank "main/method_example"
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

	fmt.Println("============ Slice ==================")
	array := [5]int{1, 2, 3, 4, 5}
	slice := array[1:2]

	fmt.Println("array:", array)
	fmt.Println("slice:", slice, len(slice), cap(slice))

	array[1] = 100

	fmt.Println("array:", array)
	fmt.Println("slice:", slice, len(slice), cap(slice))

	account := bank.Account{}
	account1 := &account
	fmt.Printf("account1 ::: %p", account1)
	fmt.Printf("account ::: %p", account)
	account.Deposit(100)
	account1.WithDraw(50)
	fmt.Println("1 ::: ", account.GetBalance())
	fmt.Println("2 ::: ", account1.GetBalance())
	account.WithDraw(20)
	fmt.Println("3 ::: ", account.GetBalance())
	fmt.Println("4 ::: ", account1.GetBalance())
	balance := account.WithDrawByValue(30)
	fmt.Println("5 ::: ", account.GetBalance())
	fmt.Println("6 ::: ", account1.GetBalance())
	fmt.Println("7 ::: ", balance)
	bank.WithDrawFunc(&account, 30)
	fmt.Println(account.GetBalance())

	newAccount := bank.NewAccount("id", "black")
	newAccount.Deposit(1000)
	fmt.Println(newAccount.GetBalance())
}
