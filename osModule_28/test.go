package main

import (
	"fmt"
	"os"
)

func test1()  {
	file, _ := os.Open("a.txt")
	//创建buffer缓冲区
	buf:=make([]byte,5)
	//读取文件到缓冲区
	count, _ := file.ReadAt(buf, 3)  //往后移3位读取5长度文件
	//打印文件内容
	fmt.Println(string(buf))
	//打印字节数
	fmt.Println(count)
}
func test2()  {
	file, _ := os.Open("a.txt")
	_, _ = file.Seek(3, 0) //往后移3位读取5长度文件
	buf:=make([]byte,5)
	count, _ := file.Read(buf)
	//打印文件内容
	fmt.Println(string(buf))
	//打印字节数
	fmt.Println(count)
}
//func main()  {
//	test1()
//    test2()
//}