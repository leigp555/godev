package main

import (
	"fmt"
	"sync"
)

var num int =100
var wg sync.WaitGroup
var lock sync.Mutex
func add()  {
	lock.Lock()
	num+=1
	wg.Done()
	lock.Unlock()
}
func sub()  {
	lock.Lock()
	num-=1
	wg.Done()
	lock.Unlock()
}
func test1()  {
	for i := 0; i < 100; i++ {
		go add()
		wg.Add(1)
		go sub()
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(num)
}
func main()  {
  test1()
}
