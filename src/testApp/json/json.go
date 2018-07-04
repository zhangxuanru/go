package main

import (
	"encoding/json"
	"fmt"
)

//传统方式 解析JSON
func main()  {
	decodeJson()
    fmt.Println("\n")
	decodeInterfaceJson();
}

//注意大小写 即首字母大写 知道数据结构时采用这种方法
type Server struct {
	ServerName string
	ServerIP string
}

//知道数据结构时采用这种方法
type Serverslice struct {
	   Servers []Server   //注意大小写 即首字母大写
}

func decodeJson()  {
	var s  Serverslice
	jsonStr := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	json.Unmarshal([]byte(jsonStr),&s)
	fmt.Println(s)
}


//不知道数据结构时采用这种方法
func decodeInterfaceJson()  {
	var f interface{}
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	err := json.Unmarshal(b,&f)
	if err!=nil{
		fmt.Println(err)
	}

	//存到interface中，下面将用这种方式来访问
	m := f.(map[string]interface{})
	for k,v := range m{
		  fmt.Println(k,":",v)
	}
    fmt.Println("\n")

	for k,v := range m{
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case float64:
			fmt.Println(k,"is float64",vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}



	fmt.Println(f,m)
}


