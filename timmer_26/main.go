package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
func clock(duration time.Duration)  {
	timer:=time.NewTimer(duration)
	<-timer.C
}
func test1()  {
	fmt.Println(time.Now())
	clock(time.Second*5)
	fmt.Println(time.Now())
	fmt.Println("main end ...")
}
func test2()  {
	fmt.Println(time.Now())
	time.Sleep(time.Second*5)
	fmt.Println(time.Now())

}
func test3()  {
	fmt.Println(time.Now())
	<-time.After(time.Second*5)
	fmt.Println(time.Now())
}
func test4()  {
	timer:=time.NewTimer(time.Second*3)
	wg.Add(1)
	go func() {
		println("延迟3秒执行")
		<-timer.C
		wg.Done()
	}()
	wg.Wait()
}
func test5()  {
	timer:=time.NewTimer(time.Second*3)
	go func() {
		<-timer.C
		fmt.Println("延迟3秒执行")
	}()
	stop := timer.Stop()
	if stop{
		fmt.Println("stop...")
	}
}
func test6()  {
	timer:=time.NewTicker(time.Second*1)
	count:=0
	for _ = range timer.C {
		fmt.Println("ticker")
		count+=1
		if count>5{
			timer.Stop()
			break
		}
	}
}
func main()  {
	defer fmt.Println("main end ...")
	//test1()
	//test2()
	//test3()
	//test4()
	//test5()
	//time.Sleep(time.Second*5)
	test6()
}
