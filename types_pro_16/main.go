package main

import "fmt"

func test1()  {
	type myInt int
	var x myInt=100
	fmt.Printf("%v\n",x)
	fmt.Printf("%T\n",x)
}
func test2()  {
	type myInt =int
	var y myInt=200
	fmt.Printf("%v\n",y)
	fmt.Printf("%T\n",y)
}
func main()  {
	test1()
	test2()
}
