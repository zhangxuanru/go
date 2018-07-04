package main

/*
https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/05.6.md

https://github.com/astaxie/goredis
redis 操作
*/
import (
	"github.com/astaxie/goredis"
	"fmt"
)

var client goredis.Client

func main()  {
	client.Addr = "127.0.0.1:6379"

	client.Set("aa",[]byte("world"))

	val,_ := client.Get("aa")

	fmt.Println(string(val))

	client.Del("aa")


	vals := []string{"a","b","c","d","e"}
	for _,v := range vals{
		client.Rpush("list",[]byte(v))
	}

	detailList,_ := client.Lrange("list",0,4)
	for i,v := range detailList{
          println(i,":",string(v))
	}
	client.Del("list")
}