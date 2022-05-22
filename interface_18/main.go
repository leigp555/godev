package main

import "fmt"

type USBer interface{
	read()
	write()
}

type Computer struct {
	name string
}
func (Computer) read()  {
	fmt.Println("读接口")
}
func (Computer)write()  {
	fmt.Println("写接口")
}

func test1()  {
	var device USBer=Computer{name:"苹果"}
	device.write()
	device.read()
}

type Player interface {
	 playMusic(string)string
}
type Video interface {
	playVideo(string)string
}

type Mobile struct {

}

func (Mobile)playMusic(name string) string {
	return name
}
func (Mobile)playVideo(name string) string {
	return name
}
func test2()  {
	var mobile Player=Mobile{}
	music := mobile.playMusic("明天你好")
	fmt.Println(music)
	var mobile2 Video=Mobile{}
	video := mobile2.playVideo("肖申克的救赎")
	fmt.Println(video)
}


//func main()  {
//	//test1()
//	test2()
//}
