package main

import "fmt"

func test(a int)int{
	if a==1 {
		return 1
	}else {
		return a*test(a-1)
	}
}
func test2()  {
	fmt.Println("start")
	defer fmt.Println("step1")
	defer fmt.Println("step2")
	defer fmt.Println("step3")
	fmt.Println("end...")
}

var initNum int=givNum()
var x int
func init()  {
	fmt.Println("init")
	x=100
}
func givNum()int{
	fmt.Println("initNum")
	return 100
}

func main()  {
	//ret := test(5)
	//fmt.Println(ret)
	//test2()
	fmt.Println("main")
	fmt.Println(x)
}