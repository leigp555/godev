package main

import "fmt"

type Tom struct {
	name string
	weight int
	height int
	gender string
}
type Man struct {
	name string
	age int
	children Tom
}
func test(){
	var kind= Tom{"tom",120,180,"男"}
	var bob=Man{name:"Bob",age: 32,children:kind}
	bob.children.height=175
	bob.age=40
	fmt.Println(bob)
	fmt.Println(bob.children)
}

type Customer struct {
	name string
}



func (Customer) login(name string,password int) bool {
	if name=="bob"&&password==123{
		return true
	}else {
		return false
	}
}

func trying()  {
	var custom=Customer{name:"tob"}
	isLogin := custom.login("bob", 123)
	fmt.Println(isLogin)
}

func (cus *Customer)trying2()  {
	cus.name="xxx"
}
func (cus Customer)trying3()  {
	cus.name="xxx"
}
func trying4()  {
	var custom=Customer{name: "小黑"}
	fmt.Println(custom)
	custom.trying3()
	fmt.Println(custom)
	custom.trying2()
	fmt.Println(custom)
}
func main()  {
	//test()
	//trying()
	trying4()
}