package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup
var data int32=100                  //需要指定明确的int类型
func add()  {
	atomic.AddInt32(&data,1)
	wg.Done()
}
func sub(){
	atomic.AddInt32(&data,-1)
	wg.Done()
}

func test1()  {
	for i := 0; i < 100; i++ {
		wg.Add(2)
		go add()
		go sub()
	}
	wg.Wait()
	fmt.Println(data)
}
func test2()  {
	var num int32 =150
	x := atomic.LoadInt32(&num)
	fmt.Println(x)
}
func test3()  {
	var num int32 =150
	atomic.StoreInt32(&num,280)
	fmt.Println(num)
}
func test4() {
	var num int32 =200
	b:=atomic.CompareAndSwapInt32(&num,200,500)
	fmt.Println(b)
	fmt.Println(num)
}
func main()  {
	//test1()
	//test2()
	//test3()
	test4()
}
