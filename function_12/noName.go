package main

import "fmt"

func noName(x int,y int) int {
	add:=func(a int,b int)int{
		return a+b
	}
	return add(x, y)
}
func high1() func(int,int)int {
	add:=func(a int,b int)int{
		return a+b
	}
	return add
}
func main()  {
	//ret := noName(23,33)
	//fmt.Println(ret)
	ret := high1()(23, 45)
	fmt.Println(ret)
}