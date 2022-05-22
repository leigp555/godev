package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func marshal()  {
	type Nav struct {
		//XMLNAME xml.Name `xml:"nav"`
		Main string `xml:"main"`
		Div  string `xml:"div"`
	}
	html:=Nav{Main: "主要内容", Div: "我是div"}
	marshal, _ := xml.MarshalIndent(html," "," ")
	fmt.Printf(string(marshal))
}
func unMarshal()  {
	type Nav struct {
		//XMLNAME xml.Name `xml:"nav"`
		Main string `xml:"main"`
		Div  string `xml:"div"`
	}
	s:=`
      <Nav>
        <main>主要内容</main>
        <div>我是div</div>
      </Nav>
    `
	ht:=[]byte(s)
	var sht Nav
	err := xml.Unmarshal(ht, &sht)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Printf("%v",sht)
}

func read()  {
	type Nav struct {
		//XMLNAME xml.Name `xml:"nav"`
		Main string `xml:"main"`
		Div  string `xml:"div"`
	}
	ht, _ := ioutil.ReadFile("a.xml")
	var sht Nav
	err := xml.Unmarshal(ht, &sht)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Printf("%v",sht)
}
func main()  {
	type Nav struct {
		//XMLNAME xml.Name `xml:"nav"`
		Main string `xml:"main"`
		Div  string `xml:"div"`
	}
	xmlTem:=Nav{Main: "主要内容", Div: "我是div"}
	wr, _ := os.OpenFile("a.xml", os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer wr.Close()
	encode := xml.NewEncoder(wr)
	encode.Encode(xmlTem)

}

