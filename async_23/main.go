package main

import (
	"fmt"
	"sync"
)
var wg sync.WaitGroup
func test(num int)  {

	fmt.Println(num)
	defer wg.Done()
}
func main()  {
	for i:=0;i<10;i++{
		go test(i)
		wg.Add(1)
	}
	wg.Wait()
}
