package main

/*
高并发体现: 用户的每一次请求都是在一个新的goroutine去服务，相互不影响。

与我们一般编写的http服务器不同, Go为了实现高并发和高性能, 使用了goroutines来处理Conn的读写事件, 这样每个请求都能保持独立，相互不会阻塞，可以高效的响应网络事件。这是Go高效的保证。

*/
import (
	"net/http"
	"log"
	"fmt"
	"strings"
)

func sayHello(w http.ResponseWriter, r *http.Request)  {
    r.ParseForm()
    fmt.Println(r.Form)
    fmt.Println("path:",r.URL.Path)
    fmt.Println("Scheme:",r.URL.Scheme)
    for k,v := range r.Form{
    	fmt.Println("k:",k)
    	fmt.Println("val:",strings.Join(v,""))
	}
	fmt.Fprint(w,"htllo world")
}

func main() {
	 http.HandleFunc("/",sayHello)
	 err := http.ListenAndServe(":9090",nil)
	 if err != nil{
	 	 log.Fatal("ListenAndServe: ", err)
	 }
}




