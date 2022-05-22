package main

import "fmt"

func appendTest()  {
	arr:=[]int{1,2,3}
	arr = append(arr, 4,5,6)
	fmt.Println(arr)
}
func testLen()  {
	arr:=[...]int{1,2,3,5,6}
	fmt.Println(len(arr))
}
func newTest()  {
	b:=new(bool)
	fmt.Println(*b)
}
func makeChan()  {
	arr:=make([]int,10)
	fmt.Println(arr)
	ch:=make(chan int)
	go func() {
		ch<-100
	}()
	fmt.Println(<-ch)
}
func makeMap()  {
	m:=make(map[string]string)
	m["name"]="lgp"
	m["age"]="20"
	fmt.Println(m)
}
func makeArr()  {
	arr:=make([]int,0)
	arr = append(arr, 1,2,3,4,5,6)
	fmt.Println(arr)
}
func main()  {

}
