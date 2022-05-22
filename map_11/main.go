package main

import "fmt"

func test()  {
	var m1 =map[string]string{
		"name":"lgp",
		"age":"20",
	}
	fmt.Println(m1)
}
func test2()  {
	var m =make(map[string]string)
	m["name"]="lgp"
	m["age"]="20"
	fmt.Println(m)
}
func test3()  {
	var m =make(map[string]string)
	m["name"]="lgp"
	m["age"]="20"
	fmt.Println(m["name"])
	fmt.Println(m["age"])
}
func test4()  {
	var m1 =map[string]string{
		"name":"lgp",
		"age":"20",
	}
	s ,ok:= m1["name"]
	fmt.Println(s)
	fmt.Println(ok)
}
func test5()  {
	var m=map[string]string{
		 "name":"lgp",
		 "age":"20",
	}
	for key,value:=range m{
		fmt.Println(key+":"+value)
	}
}

func main()  {
	//test()
	//test2()
	//test3()
	//test4()
	test5()
}
