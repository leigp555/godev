package main

import (
	"fmt"
	"log"
	"os"
)

func logTest()  {
	log.Print("ok")
	log.Printf("%v","ok")
	log.Println("ok")
	log.Panic("ok")
	log.Fatal("ok")
	defer fmt.Println("xx")
}

func init()  {
	log.SetFlags(log.Ldate|log.Ltime|log.Lshortfile)
	log.SetPrefix("lgp:")
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
	if err!=nil{
		log.Fatal("日志输出错误")
	}
	log.SetOutput(file)

	//log.New(file,"lgp:",log.Ldate|log.Ltime|log.Lshortfile)
}

func main()  {
	log.Println("ok")
}
