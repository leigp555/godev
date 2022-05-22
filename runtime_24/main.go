package main

import (
	"fmt"
	"runtime"
	"time"
)

func test1()  {
	for i := 0; i < 5; i++ {
		fmt.Println("java")
	}
}
func trying1()  {
	go test1()
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println("go")
	}
	fmt.Println("end")
}

func test2()  {
	for i := 0; i < 10; i++ {
		if i>5{
			runtime.Goexit()
		}
		fmt.Printf("%v:%v\n",i,"go")
	}
}
func trying2()  {
	go test2()
}
func test3()  {
	for i := 0; i < 5; i++ {
		fmt.Println("a")
	}
}
func test4()  {
	for i := 0; i < 5; i++ {
		fmt.Println("b")
	}
}
func trying4()  {
	go test3()
	go test4()
}
func main()  {
	//trying()
	//trying2()
	fmt.Println(runtime.NumCPU())
	runtime.GOMAXPROCS(12)
	trying4()
	time.Sleep(time.Second)
}
