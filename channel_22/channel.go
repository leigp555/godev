package main

import (
	"fmt"
	"math/rand"
	"time"
)


func send(ch chan int) {
	rand.Seed(time.Now().UnixNano())
	value:=rand.Intn(10)
	fmt.Printf("send:%v\n",value)
    ch <- value
}
func test1()  {
	ch:=make(chan int,10)
	go send(ch)
	value:=<-ch
	time.Sleep(time.Second)
	defer close(ch)
	fmt.Printf("get:%v\n",value)
}

var c =make(chan int)

func test2()  {
	for i := 0; i < 10; i++ {
		c<-i
	}
	defer close(c)
}
func test3()  {
	go test2()
	for  {
		value,ok:=<-c
		if ok{
			fmt.Printf("get:%v\n",value)
		}else{
			break
		}

	}
}
func test4()  {
	go test2()
	for value := range c {
		fmt.Println(value)
	}
}
//func main()  {
//    //test1()
//	test3()
//	//test4()
//}