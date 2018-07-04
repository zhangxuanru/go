package main

import (
	"encoding/xml"
	"os"
	"fmt"
	"io/ioutil"
)

type Recurlyservers struct {
	  XMLName xml.Name `xml:"servers"`
	  Version string   `xml:"version,attr"`
	  Svs     []server  `xml:"server"`
	  Description string `xml:",innerxml"` 
}

type server struct {
	   XMLName xml.Name   `xml:"server"`
       ServerName string  `xml:"serverName"`
	   ServerIP   string  `xml:"serverIP"`
}

//读取XML
func main()  {
	getXML()
}


func getXML() {
		file,err := os.Open("src/testApp/XML/test.xml")
		if err!=nil{
		fmt.Printf("error: %v", err)
		return
		}
		defer file.Close()

		data,err := ioutil.ReadAll(file)
		if err!=nil{
		fmt.Printf("error: %v", err)
		return
		}

		v := Recurlyservers{}
		err = xml.Unmarshal(data,&v)
		if err != nil{
		fmt.Printf("error: %v", err)
		return
		}
		fmt.Println(v)
}



