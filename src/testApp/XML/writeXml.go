package main

import (
	"encoding/xml"
	"os"
	"fmt"
	"io/ioutil"
)

//生成XML文件
func main()  {
	  writeXml()
}

type Servers struct {
	XMLName xml.Name  `xml:"servers"`
	Version string    `xml:"version,attr"`
	Svs     []server   `xml:"server"`
}

type server struct {
	ServerName string `xml:"serverName"`
	ServerIP   string `xml:"serverIP"`
}

func writeXml()  {
	v := &Servers{Version:"1.0"}
	v.Svs = append(v.Svs,server{ServerName:"localhost",ServerIP:"192.168.33.10"})
	v.Svs = append(v.Svs,server{"127.0.0.1","127.0.0.1"})
	output,err := xml.MarshalIndent(v," ","   ")
	if err != nil{
		fmt.Printf("error: %v\n", err)
	}

	//os.Stdout.Write([]byte(xml.Header))
	//os.Stdout.Write(output)

    xmlData := append([]byte(xml.Header),output...)
    ioutil.WriteFile("src/testApp/XML/server.xml",xmlData,os.ModeAppend)
	fmt.Println("create xml file ok")



	////加入XML头
	//headerBytes := []byte(xml.Header)
	////拼接XML头和实际XML内容
	//xmlOutPutData := append(headerBytes, output...)
	////写入文件
	//ioutil.WriteFile("server.xml", xmlOutPutData, os.ModeAppend)

}

