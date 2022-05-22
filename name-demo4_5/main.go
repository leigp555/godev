package main

import "fmt"

func main()  {
	nameDefine := "name"
	var x string ="world"
	fmt.Println(nameDefine)
	fmt.Println(x)
	const (
		a1=iota
		a2=100
		a3=iota
	)
    fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
}