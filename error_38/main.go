package main

import (
	"errors"
	"fmt"
	"log"
)

func str(s string)(string,error) {
	if s ==""{
		err := errors.New("字符串不能为空")
		return "",err
	}else {
		return s,nil
	}
}
func main()  {
	s, err := str("")
	if err!=nil{
		fmt.Println(err.Error())
		log.Fatal(err)
	}
	fmt.Println(s)
}
