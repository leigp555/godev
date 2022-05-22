package main

import "fmt"

type User struct{
	Name string
}
func main(){
	user := User{Name:"lgp"}
	fmt.Printf("%v\n",user)
	fmt.Printf("%#v\n",user)
	fmt.Printf("%T\n",user)
	str:="hello world"
	fmt.Printf("%s\n",str)
	fmt.Printf("%s\n",[]byte("hello world"))
	fmt.Printf("%q\n",str)
	fmt.Printf("%p",&str)
}