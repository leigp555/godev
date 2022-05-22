package main

import "fmt"

var chanInt =make(chan int)
var chanStr =make(chan string)

func test5()  {
	chanInt<-100
	chanStr<-"hello"
	close(chanInt)
	close(chanStr)
}
func trying5()  {
	go test5()
	for{
		select {
		case r:=<-chanInt:
			fmt.Println(r)
		case r:=<-chanStr:
			fmt.Println(r)
		default:
			fmt.Println("无值可读")
		}
	}
}
func main()  {
	trying5()
}