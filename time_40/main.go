package main

import (
	"fmt"
	"time"
)

func timeRef()  {
	now:=time.Now()
	fmt.Println(now)
	fmt.Printf("%T",now)
	now.Year()
	now.Month()
	now.Day()
	now.Hour()
	now.Minute()
	now.Second()
}
func unix()  {
	nowUnix:=time.Now().UnixNano()
	//nowUnix:=time.Now().UnixNano()
	fmt.Printf("%T\n",nowUnix)
	fmt.Println(nowUnix)
}
func unUnix()  {
	unixNow:=time.Now().Unix()
	now:=time.Unix(unixNow,0)
	fmt.Println(now)
}
func addTime()  {
	now := time.Now()
	add := now.Add(time.Hour * 24*10+time.Minute*10+time.Second*50)
	fmt.Println(add)
}
func subTime()  {
	startTime:=time.Now()
	endTime:=time.Now().Add(time.Hour*2)
	space := endTime.Sub(startTime)
	fmt.Println(space)
}
func equal()  {
	t1:=time.Now().Add(time.Hour*2)
	t2:=time.Now()
	b := t1.Equal(t2)
	fmt.Println(b)
}
func after()  {
	t1:=time.Now().Add(time.Hour*2)
	t2:=time.Now()
	b := t1.After(t2)
	fmt.Println(b)
}
func before()  {
	t1:=time.Now().Add(time.Hour*2)
	t2:=time.Now()
	b := t1.Before(t2)
	fmt.Println(b)
}
func ticker()  {
	ch := time.Tick(time.Second)
	for i:=range ch {
		fmt.Println(i)
	}
	//ch := time.NewTicker(time.Second*2)
	//for i:=range ch.C{
	//	fmt.Println(i)
	//}
}
func format()  {
	now := time.Now()
	formatTime := now.Format("2006年01月02日15时04分")
	fmt.Println(formatTime)
}
func main()  {
	//加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err!=nil{
		fmt.Println(err)
		return
	}
	timeObj,_:=time.ParseInLocation("2006/01/02 15:04:05","2022/05/22 15:56:30",loc)
	fmt.Println(timeObj)
}

