package main

import "fmt"

func test1()  {
	var arr1 [3]int
	var arr2 [2]string
	fmt.Printf("%T\n",arr1)
	fmt.Printf("%T\n",arr2)
}
func test2()  {
	var arr1 [3]int=[3]int{2,3,5}
	var arr2=[2]string{"hello","world"}
	var arr3=[...]int{3,4,5,6,7}
	var arr4=[...]int{0:1,2:2}
	fmt.Println(arr1)
	fmt.Println(arr2)
	//fmt.Printf("%T",arr3)
	fmt.Printf("%v\n",len(arr3))
	fmt.Println(arr4)
}
func test3()  {
	var arr=[...]int{1,3,4,5,6}
	arr[1]=100
	fmt.Println(arr)
}
func test4()  {
	var arr =[...]int{1,2,3,4,5}
	fmt.Println(arr[1])
}
func test5()  {
	var arr =[...]int{1,3,4,5,6,8,5}
	for i:=0;i<len(arr);i++{
		fmt.Println(arr[i])
	}
}
func test6()  {
	var arr=[...]int{1,3,4,5,7}
	for index, value := range arr {
		fmt.Printf("%v-%v\n",index,value)
	}
}
func test7()  {
	var myNumber []int
	fmt.Printf("%T\n",myNumber)
}
func test8()  {
	var arr []int
	var count = make([]int,2)
	arr = append(arr, 100)
	arr = append(arr, 50)
	count = append(count, 200)
	count = append(count, 40)
	fmt.Println(arr)
	fmt.Println(count)
	fmt.Println(len(arr))
	fmt.Println(len(count))
	fmt.Println(cap(arr))
	fmt.Println(cap(count))
}
func test9()  {
	var arr =[]int{1,2,3,4,5,6,7,8}
	s1:=arr[0:3]
	fmt.Println(s1)
}
func test10 (){
	var arr =[]int{1,2,3,44,5}
	for index,value:=range arr{
		fmt.Printf("%v:%v\n",index,value)
	}
}
func test11 (){
	var arr =[]int{1,2,3,44,5}
	for i:=0;i<len(arr);i++{
		fmt.Println(arr[i])
	}
}
func test12 (){
	var arr []int
	arr=append(arr,1,2,3,4)
	fmt.Println(arr)
}
func test13 (){
	var arr =[]int{1,2,3,44,5}
	arr =append(arr[:2],arr[3:]...)
	fmt.Println(arr)
}
func test14 (){
	var arr =[]int{1,2,3,44,5}
	arr[1]=100
	fmt.Println(arr)
}
func test15 (){
	var arr =[]int{1,2,3,44,5}
	//浅拷贝
	var arr5 =arr         //arr2与arr的内存地址一样一个修改另一个也会修改
	//深拷贝
	var arr2=make([]int,len(arr))         //先要make一下拿到新的内存地址
	copy(arr2,arr)                         //将arr的数组拷贝到arr2中
	//删除
	arr=append(arr[:2],arr[3:]...)
	arr[1]=200
	fmt.Println(arr)
	fmt.Println(arr2)
	fmt.Println(arr5)
}
func main()  {
	//test1()
	//test2()
	//test3()
	//test4()
	//test5()
	//test6()
	//test7()
	//test8()
	//test9()
	//test10()
	//test11()
	//test12()
	//test13()
	//test14()
	test15()
}
