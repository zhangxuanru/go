package main

import "fmt"
import (
	"encoding/json"
	"os"
)

/**
输出JSON

https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/07.2.md

我们看到上面的输出字段名的首字母都是大写的，如果你想用小写的首字母怎么办呢？把结构体的字段名改成首字母小写的？JSON输出的时候必须注意，只有导出的字段才会被输出，如果修改字段名，那么就会发现什么都不会输出，所以必须通过struct tag定义来实现：
 */
func main()  {
	outputJson()
	outputJson2()
}

/*
1,要导出的字段首字母大写,
type Server struct {
	ServerName string
	ServerIP   string
}

type Serverslice struct {
    Data []Server
}

2.如果不想字段首字母大写则要采用struct tag定义来实现
*/
type Server struct {
	ServerName string `json:"serverName"`
	ServerIP   string `json:"serverIP"`
}

type Serverslice struct {
    Data []Server   `json:"data"`
}

func outputJson()  {
     var s Serverslice
     s.Data = append(s.Data,Server{ServerName:"localhost",ServerIP:"127.0.0.1"})
     s.Data = append(s.Data,Server{"www.baidu.com","23.56.77.19"})
     v,err := json.Marshal(s)
	 if err != nil {
		fmt.Println("json err:", err)
	 }
	 fmt.Println(string(v))
}

//struct tag
func outputJson2()  {
	type Server struct {
		// ID 不会导出到JSON中
		ID int `json:"-"`
		// ServerName2 的值会进行二次JSON编码
		ServerName  string `json:"serverName"`
		ServerName2 string `json:"serverName2,string"`
		// 如果 ServerIP 为空，则不输出到JSON串中
		ServerIP   string `json:"serverIP,omitempty"`
	}

	s := Server {
		ID:         3,
		ServerName:  `Go "1.0" `,
		ServerName2: `Go "1.0" `,
		ServerIP:   ``,
	}
	b, _ := json.Marshal(s)
	os.Stdout.Write(b)
}