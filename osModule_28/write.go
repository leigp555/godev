package main

import (
	"fmt"
	"os"
	"time"
)

func writeF()  {
	//file, _ := os.OpenFile("a.txt", os.O_RDWR|os.O_APPEND, 0755)
	//write, err := file.Write([]byte("hello world"))

	//writeString, err2 := file.WriteString("hello world")

	//at, err2 := file.WriteAt([]byte("hello world"), 3)
	//err := file.Close()
}
func newProcess(){
	//获取进程id
	fmt.Println(os.Getegid())
	//获取父进程id
	fmt.Println(os.Getppid())

	//设置新进程的属性
	attr:=&os.ProcAttr{
		//指定新进程继承的活动文件对象
		//前三个分别为标准输入，标准输出，标准错误输出
		Files: []*os.File{os.Stdin,os.Stdout,os.Stderr},
		//新进程的环境变量
		Env:os.Environ(),
	}
	//开始一个新的进程
	p,err:=os.StartProcess("C:\\Windows\\System32\\notepad.exe",[]string{"C:\\Windows\\System32\\notepad.exe","D:\\a.txt"},attr)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(p)
	fmt.Println("进程id",p.Pid)
	//通过进程id查找进程
	p2, _ := os.FindProcess(p.Pid)
	fmt.Println(p2)

	//等待10秒执行函数
	time.AfterFunc(time.Second*10, func() {
		//想进程发送退出信号
		err2 := p.Signal(os.Kill)
		if err2!=nil{
			fmt.Println(err2)
		}
	})
	//等待进程p的退出，返回进程状态
	wait, _ := p.Wait()     //这里是阻塞的一直在等待进程的退出
	fmt.Println(wait.String())
}

//func main()  {
//	newProcess()
//}