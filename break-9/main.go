package main

import "fmt"

func test()  {
	for i:=0;i<10;i++{
        fmt.Println(i)
		if i>5{
			break
		}
	}
}
func test2()  {
	i :=2
	switch i {
	case 1:
		fmt.Println(i)
		break
	case 2:
		fmt.Println(i)
		break
		fallthrough
	case 3:
		fmt.Println(i)
	}
}
func test3()  {
	MYLABEL:
	for i:=0;i<10;i++{
		fmt.Println(i)
		if i>5{
			break MYLABEL
		}
	}
}
func test4()  {
	for i:=0;i<10;i++{
		if i%2==0{
			fmt.Println(i)
		}else {
			continue
		}
	}
}
func test5()  {
MYLABEL:
	for i:=0;i<10;i++{
		for j:=0;j<10;j++ {
			if i%2!=0&&j%2==0{
				goto MYLABEL
			}
			fmt.Printf("%v,%v\n",i,j)
		}
	}
	fmt.Println("end....end")
}
func test6()  {
	i:=1
	if i>2{
		fmt.Println("i")
	}else {
		goto END
	}
	END:
		fmt.Println("end..")
}
func main()  {
	//test()
	//test2()
	//test3()
	//test4()
	test5()
	//test6()
}
