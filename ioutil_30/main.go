package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func test1()  {
	r := strings.NewReader("hello world")
	buf:=make([]byte,10)
	_,err:= r.Read(buf)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println(string(buf))
}
func copyTest()  {
	reader := strings.NewReader("hello world")
	file, _ := os.OpenFile("a.tex", os.O_APPEND|os.O_RDWR|os.O_CREATE, os.ModePerm)
	_, err := io.Copy(file, reader)
	if err!=nil{
		log.Fatal(err)
	}

}
func main() {
   copyTest()
}