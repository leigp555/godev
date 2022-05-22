package main

import "fmt"

func sun(a int ,b int) (ret int) {
	ret=a+b
	return ret
}
func sun2(a int ,b int) (x int , y int , ret int) {
	ret=a+b
	return  a,b,ret
}
func user() (name string,age int) {
	name="lgp"
	age=20
	return
}
func user2() ( string, int) {
	name:="lgp"
	age:=20
	return name,age
}

func test(s []int)  {
	s[1]=100
}
func test2(args ...int)  {
	for _, i2 := range args {
		fmt.Println(i2)
	}
}

type fun func(int,int)int

func add(a int,b int) int {
	return a+b
}
func max(a int,b int) int {
	if a>b {
		return a
	}else {
		return b
	}
}

func high(name string,f func(string))  {
	f(name)
}
func printName(content string)  {
	fmt.Println(content)
}

func add1(a int,b int)int  {
	return a+b
}
func sub1(a int,b int)int  {
	return a-b
}
func calc(action string)func(int,int)int {
	switch action {
	case "+":
		return add1
	case "-":
		return sub1
	default:
		return nil
	}
}

//func main()  {
	//ret := sun(2, 4)
	//fmt.Println(ret)

	//x, y, ret := sun2(4,2)
	//fmt.Printf("%v:%v:%v\n",x,y,ret)

	//name, age := user()
	//fmt.Printf("%v-%v",name,age)

	//name, age := user2()
	//fmt.Printf("%v-%v",name,age)

	//var arr=[]int{1,2,3,4}
	//test(arr)
	//fmt.Println(arr)

	//test2(1,2,3,4,5,6,7)

	//var f fun
	//f=add
	//ret := f(5, 2)
	//fmt.Println(ret)
	//f=max
	//i := f(5, 20)
	//fmt.Println(i)


	//high("lgp",printName)
//	ret := calc("-")(5, 2)
//	fmt.Println(ret)
//}