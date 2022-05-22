package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Person struct {
	Name string
	Age int
	Gender string
}

func toJson()  {
	tom :=Person{Name:"lgp",Age:20,Gender:"男"}
	t, _ := json.Marshal(tom)
	fmt.Printf("%v",string(t))
}
func toStruct()  {
	arr:=[]byte(`{"Name":"lgp","Age":20,"Gender":"男"}`)
	var p Person
	err := json.Unmarshal(arr, &p)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Printf("%v",p)

	//arr:=[]byte(`{"Name":"lgp","Age":20,"Gender":"男"}`)
	//var p interface{}
	//err := json.Unmarshal(arr, &p)
	//if err!=nil{
	//	log.Fatal(err)
	//}
	//fmt.Printf("%T",p)
}
func forMap()  {
	m:=make(map[string]string)
	m["name"]="lgp"
	m["age"]="20"
	for s, s2 := range m {
		fmt.Printf("%v:%v\n",s,s2)
	}
}
func forJson()  {
	arr:=[]byte(`{"Name":"lgp","Age":20,"Gender":"男"}`)
	//var p interface{}
	var p map[string]interface{}
	err := json.Unmarshal(arr, &p)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Printf("%v\n",p)
	for s, i := range p {
		fmt.Printf("%v:%v\n",s,i)
	}
}
func decode()  {
	file, _ := os.OpenFile("a.json", os.O_RDWR, os.ModePerm)
	defer file.Close()
	decode := json.NewDecoder(file)
	var v map[string]interface{}
	//var v Person
	decode.Decode(&v)
	fmt.Println(v)
}
func main()  {
	type Pet struct {
		Name string
		Color string
		Children []string
	}
	cat:=Pet{
		Name: "大橘",
		Color: "orange",
		Children: []string{"中橘","黑橘","小橘"},
	}
	file, _ := os.OpenFile("a.json", os.O_RDWR|os.O_APPEND, os.ModePerm)
	defer file.Close()
	encode := json.NewEncoder(file)
	encode.Encode(cat)
}
