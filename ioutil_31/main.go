package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func testReadAll()  {
	r, _ := os.OpenFile("a.txt", os.O_RDONLY, os.ModePerm)

	defer r.Close()
	all, err := io.ReadAll(r)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println(string(all))
}

func readDir()  {
	dir, err := ioutil.ReadDir(".")
	if err!=nil{
		log.Fatal(err)
	}
	for _,file := range dir {
		fmt.Println(file.Name())
	}
}
func readFile()  {
	file, err := ioutil.ReadFile("a.txt")
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println(string(file))
}
func writeFile()  {
	buf:=[]byte("hello world people")
	err := ioutil.WriteFile("a.txt", buf, os.ModePerm)
	if err!=nil {
		log.Fatal(err)
	}
}
func main()  {

}
