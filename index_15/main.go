package main

import "fmt"

func test(x *string)  {
	*x="lgp"
}

func test2()  {
	var ip *int
	var x int=100
	ip=&x
	fmt.Printf("%v\n",ip)
	fmt.Printf("%T\n",ip)

}
func test3()  {
	arr:=[3]int{1,3,5}
	var ip [3]*int
	for i:=0;i<len(arr);i++{
		ip[i]=&arr[i]
	}
	fmt.Printf("%v\n",ip)

	for i:=0;i<len(ip);i++{
		fmt.Printf("%v\n",*ip[i])
	}
}
func main()  {
	//s:="hello"
	//test(&s)
	//fmt.Println(s)
	//test2()
	test3()
}
