package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"regexp"
)

func trans()  {
	var a int=10
	var b byte=100
	var c string="hello"
	var d []byte=[]byte("world")
	//数字转字节
	x:=byte(a)
	fmt.Println(x)
	fmt.Println("n----b")
	//字节转数字
	y:=int(b)
	fmt.Println(y)
	fmt.Println("b----n")
	//字符串转字节切片
	z := []byte(c)
	fmt.Println(z)
	fmt.Println("s----b[]")
	//字节切片转字符串
	s:=string(d)
	fmt.Println(s)
	fmt.Println("b[]----s")
}
func contain()  {
	s:="hello world"
	b:=[]byte(s)
	x:=[]byte("hello")
	y:=[]byte("WORLD")
	have := bytes.Contains(b, x)
	have2 := bytes.Contains(b, y)
	fmt.Println(have)
	fmt.Println(have2)
}
func count(){
	by:=[]byte("hhhhhllllllooooowwwwwwwwwwwww")
	x:=[]byte("h")
	y:=[]byte("l")
	z:=[]byte("w")
	fmt.Println(bytes.Count(by,x))
	fmt.Println(bytes.Count(by,y))
	fmt.Println(bytes.Count(by,z))
}
func repeat()  {
	by:=[]byte("ok")
	nby := bytes.Repeat(by, 10)
	fmt.Println(string(nby))
}
func replace()  {
	by:=[]byte("ok hsfsf ok fhgdg ok")
	by1:=[]byte("ok")
	by2:=[]byte("no")
	nby := bytes.Replace(by, by1, by2, -1)
	fmt.Println(string(nby))
    ////正则
	//text :="abcgodef go good"
	//reg := regexp.MustCompile(`\bgo\b`)
	//fmt.Printf("%s\n", reg.ReplaceAllString(text, "**"))
	//正则
	text :=[]byte("abcgodef go good")
	reg := regexp.MustCompile(`\bgo\b`)
	fmt.Printf("%s\n", reg.ReplaceAllString(string(text), "**"))
}
func runes() {
	by:=[]byte("你好我的世界")
	by2:=bytes.Runes(by)
	fmt.Println(len(by2))
}
func joinTest()  {
	barr:=[][]byte{[]byte("你好"),[]byte("世界")}
	mid:=[]byte(",")
	ret:=bytes.Join(barr,mid)
	fmt.Println(string(ret))
}
func byRead()  {
	re := bytes.NewReader([]byte("gdoj,hfhf,"))
	re2 := bufio.NewReader(re)
	ret, err := re2.ReadBytes(',')
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println(string(ret))
}
func readBuf() {
	re := bytes.NewReader([]byte("gdojhfhf,"))
	re2 := bufio.NewReader(re)
	buf:=make([]byte,5)
	for{
		n, err := re2.Read(buf)
		if err!=nil{
			break
		}
		fmt.Println(string(buf[0:n]))
	}
}

func bytesTest()  {
	//var b bytes.Buffer
	//b:=new(bytes.Buffer)
	//b:=bytes.NewBuffer([]byte("hello"))
	b:=bytes.NewBufferString("hello")
	b.WriteString("hello")
	s := string(b.Bytes())
	fmt.Println(s)
}
func main()  {
	file, _ := os.OpenFile("a.txt", os.O_CREATE|os.O_RDWR, os.ModePerm)
	r1 := bufio.NewReader(file)
	ch := make([]byte, 5)
	for{
		n, err := r1.Read(ch)
		if err!=nil{
			break
		}
		fmt.Println(string(ch[0:n]))
	}
}
