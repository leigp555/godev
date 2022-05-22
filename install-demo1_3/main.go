package main

import (
	"fmt"
	"installDemo/user"
)

func main() {
	fmt.Printf("\"hello\": %v\n", "hello")
	var hello string=user.Hello()
	fmt.Printf("%T",hello)
}

