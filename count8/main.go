package main

import "fmt"

func numberType(num int)string{
	if num%2==0{
		return "偶数"
	}else{
		return "奇数"
	}
}

func main()  {
	//a:=100
	//if a>=100{
	//	fmt.Println("a大于等于100")
	//}else{
	//	fmt.Println("a小于100")
	//}
	//switch a{
	//case 100:
	//	fmt.Println("a等于100")
	//case 200:
	//	fmt.Println("a等于200")
	//default:
	//	fmt.Println("a是一个数")
	//}
	////判断奇偶
	//var numberInput int
	//fmt.Println("请输入一个数")
	//_, err := fmt.Scan(&numberInput)
	//if err != nil {
	//	return
	//}
	//fmt.Println(numberType(numberInput))
	//
	//var workDay int
	//fmt.Println("今天星期几？")
	//_,err=fmt.Scan(&workDay)
	//if err != nil{
	//	return
	//}
	//switch workDay{
	//case 1,2,3,4,5:
	//	fmt.Println("工作日")
	//case 6,7:
	//	fmt.Println("休息日")
	//}
	//i :=0
	//for;i<10;i++{
	//	fmt.Printf("i= %v",i)
	//}
	//i :=0
	//for {
	//	fmt.Printf("i=%v",i)
	//	i++
	//	if i>=100{
	//		break
	//	}
	//}
	var s string="hello world"
	for index,value:=range s{
		fmt.Printf("index=%v,value=%c\n",index,value)
	}
}
