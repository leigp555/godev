package main

import "fmt"

type Fly interface {
	fly()
}
type Swim interface {
	swim()
}
type FlyFish interface {
	Fly
	Swim
}
type Fish struct {

}

func (Fish)fly()  {
	fmt.Println("fly")
}
func (Fish)swim()  {
	fmt.Println("swim")
}

//func main()  {
//	var fish FlyFish =new(Fish)
//	fish.swim()
//	fish.fly()
//}