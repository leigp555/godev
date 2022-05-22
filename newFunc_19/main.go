package main

import "fmt"

type Person struct {
	name string
	age int
}

func (Person)speak(content string)  {
	fmt.Println(content)
}

func newPerson(name string,age int)(*Person,error)  {
	if name ==""{
		return nil,fmt.Errorf("name不能为空")
	}else if age<0{
		return nil,fmt.Errorf("age不能小于0")
	}else{
		return &Person{name: name,age: age},nil
	}
}
func main()  {
	person, err := newPerson("lgp", 10)
	if err!=nil {
		fmt.Println(err)
		return
	}
	fmt.Println(*person)
	person.speak("how are you")
}
