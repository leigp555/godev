package main

import (
	"fmt"
	"sort"
)

func sortInt()  {
	x:=sort.IntSlice{4,2,5,8,1}
	sort.Ints(x)
	fmt.Println(x)
}

func sortFloat()  {
	x:=sort.Float64Slice{4.1,2.2,5.3,8.4,1.5}
	sort.Float64s(x)
	fmt.Println(x)
}
func strNum()  {
	x:=[]string{"34","23","76","43"}
	sort.Strings(x)
	fmt.Println(x)
}
func strSy()  {
	x:=[]string{"f","s","j","a"}
	sort.Strings(x)
	fmt.Println(x)
}
func sortChinese()  {
	x:=sort.StringSlice{"你","我","额","啊"}
	sort.Strings(x)
	fmt.Println(x)
	for _, v := range x {
		fmt.Println([]byte(v))
	}
}
type complex [][]int

func (l complex)Len()int{
	return len(l)
}
func (l complex)Swap(i,j int)  {
	l[i],l[j]=l[j],l[i]
}
func (l complex)Less(i,j int) bool {
	return l[i][1]>l[j][1]
}
func complexText()  {
	n:=complex{{1,2},{2,3},{5 ,6}}
	sort.Sort(complex(n))
	fmt.Println(n)            //[1 3 5 7 9]
}

type NewInts [] uint

func (n NewInts)Len()int  {
	return len(n)
}
func (n NewInts)Less(i int,j int)bool  {
	return n[i]<n[j]
}
func (n NewInts)Swap(i int , j int)  {
	n[i],n[j]=n[j],n[i]
}
func customSort()  {
	n:=NewInts{9,5,7,3,1}
	sort.Sort(NewInts(n))
	fmt.Println(n)

}

type newMap []map[string]float64
func (n newMap)Len()int  {
	return len(n)
}
func (n newMap)Less(i int,j int)bool  {
	return n[i]["a"]<n[j]["a"]
}
func (n newMap)Swap(i int , j int)  {
	n[i],n[j]=n[j],n[i]
}
func testNewMap()  {
	arr:=newMap{
		{"a":10,"b":65},
		{"a":4,"b":76},
		{"a":6,"b":23},
		{"a":1,"b":46},
	}
	sort.Sort(newMap(arr))
	fmt.Println(arr)
}

type People struct {
	name string
	age int
}
type newStruct []People
func (n newStruct)Len()int  {
	return len(n)
}
func (n newStruct)Less(i int,j int)bool  {
	return n[i].age<n[j].age
}
func (n newStruct)Swap(i int , j int)  {
	n[i],n[j]=n[j],n[i]
}
func testNewStruct()  {
	arr:=newStruct{
		{name: "lgp" ,age: 18},
		{name: "xgd" ,age: 24},
		{name: "das" ,age: 10},
		{name: "gdg" ,age: 5},
	}
	sort.Sort(newStruct(arr))
	fmt.Println(arr)
}

func main()  {
	testNewStruct()
}
