package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init()  {
	rand.Seed(time.Now().UnixMicro())
}
func randGenerate()  {
	for i := 0; i <10 ; i++ {
		fmt.Printf("%v\n",rand.Intn(100))
	}
}
func main()  {
	var a float32=3.141592653
	fmt.Println(a)
	b:=int(a)
	fmt.Println(b)
}
