package main

import "fmt"

type Pet interface {
	eat()
	sleep()
}

type Mouse struct {

}
type Iron struct {

}
type Man struct {

}
func (Mouse)eat()  {
	fmt.Println("mouse eat")
}
func (Mouse)sleep()  {
	fmt.Println("mouse sleep")
}
func (Iron)eat()  {
	fmt.Println("iron eat")
}
func (Iron)sleep(){
	fmt.Println("iron sleep")
}
func (Man)care(pet Pet)  {
	pet.sleep()
	pet.eat()
}
func main()  {
	var mouse Pet=new(Mouse)
	var iron Pet =new(Iron)
	man :=Man{}
	man.care(mouse)
	man.care(iron)
}