package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func test1()  {
	//返回一个Reader指针
	r:=strings.NewReader("helloadsfs")
	//io.copy接收一个writer和一个reader
	//将r的值赋值到os.Stdout
	_, err := io.Copy(os.Stdout, r)
	if err !=nil{
		log.Fatal(err)
	}

}
func main()  {
   test1()
}
