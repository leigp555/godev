package main

import "fmt"

func add() func(int)int {
	var x int
	return func(y int) int{
		x+=y
		return x
	}
}

func main()  {
	f:=add()
	i := f(10)
	fmt.Println(i)
	i=f(20)
	fmt.Println(i)
	i=f(30)
	fmt.Println(i)
}
