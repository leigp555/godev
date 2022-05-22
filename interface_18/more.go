package main

import "fmt"

type Eater interface {
	eat()string
}
type Cat struct {

}
type Dog struct {

}

func (Cat)eat() string {
	fmt.Println("cat eat")
	return "cat"
}
func (Dog)eat() string {
	fmt.Println("dog eat")
	return "dog"
}

//func main()  {
//	var cat Eater=new(Cat)
//	var dog Eater=new(Dog)
//	cat.eat()
//	dog.eat()
//	fmt.Printf("%p\n",cat)
//	fmt.Printf("%p\n",dog)
//}