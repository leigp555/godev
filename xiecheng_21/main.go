package main

import (
	"fmt"
	"time"
)

func test(name string)  {
	for i:=0;i<5;i++{
		fmt.Println(name)
	}
}

func main()  {
	go test("java")
	test("go")
	time.Sleep(time.Second*2)
	fmt.Println("end...")
}