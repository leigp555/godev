package main

import "fmt"

type Person struct {
	name string
	id,age int
	email string
}

func test1()  {
	var man Person
	man.id=1
	man.age=18
	man.name="lgp"
	man.email="xx@qq.com"
	fmt.Printf("%T\n",man)
	fmt.Printf("%v\n",man)
	fmt.Println(man.name)
	fmt.Println(man.email)
}
func test2()  {
	var person struct{
		name string
		age int
	}
	person.age=20
	person.name="lgp"
	fmt.Println(person)
}

type Dog struct {
	name string
	age int
	color string
}

func test3()  {
	var dog =Dog{
		name:"小白",
		age: 10,
		color: "black",
	}
	fmt.Println(dog)
}
func test4()  {
	var dog Dog
	dog.name="大白"
	dog.color="white"
	dog.age=20
	fmt.Println(dog)
}
func test5()  {
	var dog =Dog{"小黑", 20, "blue"}
	fmt.Println(dog)
}
func test6()  {
	dog:=Dog{name:"多米", color: "yellow"}
	fmt.Println(dog)
}
func test7()  {
	dog:=Dog{"小黑", 20, "blue"}
	var d *Dog
	d=&dog
	fmt.Println(dog)
	fmt.Printf("%p\n",d)
	fmt.Printf("%v\n",*d)
}
func test8()  {
	var dog = new(Dog)
	dog.name="小黑"
	dog.age=15
	dog.color="black"
	fmt.Printf("%T\n",dog)
	fmt.Printf("%p\n",dog)
	fmt.Printf("%v\n",*dog)
}
func test9(person Dog)  {
	person.name="小白"
	person.age=11
	person.color="white"
}
func test10()  {
	var dog=Dog{"小黑",20,"black"}
	test9(dog)
	fmt.Println(dog)
}


//func main()  {
	//test1()
	//test2()
	//test3()
    //test4()
	//test5()
	//test6()
	//test7()
	//test8()
//	test10()
//}