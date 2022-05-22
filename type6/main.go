package main

import (
	"bytes"
	"fmt"
	"math"
	"strings"
	"unsafe"
)

func main()  {
	stringName :="hello"
	fmt.Println(stringName)
	var arr [2]int =[2]int{1,2}
	fmt.Printf("%T\n",arr)
	var age int = 18
	gender :="男"
	if age >=18 && gender=="男"{
		fmt.Println("成年男性")
	}else{
		fmt.Println("未成年")
	}
	var count int =10
	for i:=0;i<count;i++{
		fmt.Printf("%v\n",i)
	}
	var i8 int8
	fmt.Printf("%T %dB %v~%v\n",i8,unsafe.Sizeof(i8),math.MinInt8,math.MaxInt8)
	number:=10
	fmt.Printf("%d\n",number)  //十进制输出  //10
	fmt.Printf("%b\n",number)  //二进制输出  //1010
	fmt.Printf("%o\n",number)  //八进制输出  //
	fmt.Printf("%x\n",number)  //十六进制输出  //
	fmt.Printf("%X\n",number)  //十六进制输出  //
	fmt.Printf("%f\n",math.Pi)  //浮点型输出   //A
	fmt.Printf("%.8f\n",math.Pi)  //浮点型输出   //A
	s1 := "hello"
	s2 := "world"
	//字符串链接
	s3 := s1 + s2      //方法一
	fmt.Println(s3)
	fmt.Printf("%v\n",fmt.Sprintf("%s,%s",s1,s2)) //方法二
	msg := strings.Join([]string{s1,s2},",")                       //方法三
	fmt.Println(msg)
	var buffer bytes.Buffer
	buffer.WriteString("hello")
	buffer.WriteString(",")
	buffer.WriteString("world")
	fmt.Println(buffer.String())
	str:="hello world"
	fmt.Println(strings.Split(str,""))
	fmt.Println(strings.Contains(str,"hello"))
}

