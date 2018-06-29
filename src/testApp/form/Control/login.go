package main

import (
	"net/http"
	"html/template"
	"fmt"
	"log"
	"time"
	"crypto/md5"
	"io"
	"strconv"
)

const HTMLDIR = "src/testApp/form/tpl"


/*
一个简单的登录页面
*/
func login(w http.ResponseWriter, r *http.Request)  {
	     fmt.Println("method:",r.Method)
         if r.Method == "GET" {
         	crutTime := time.Now().Unix()
         	h := md5.New()
            io.WriteString(h,strconv.FormatInt(crutTime,10))
         	token := fmt.Sprintf("%x",h.Sum(nil))  //生成md5token

            t,_ := template.ParseFiles(HTMLDIR+"/login.gtpl")
            t.Execute(w,token)
           // log.Println(t.Execute(w,nil))
		 }
		 if r.Method == "POST" {
                r.ParseForm()
                fmt.Println("username",r.Form["username"])
                fmt.Println("password",r.Form["password"])
                fmt.Println("username",r.FormValue("username"))
			    fmt.Println("password",r.FormValue("password"))
                uStr := fmt.Sprintf("username:%s, password:%s",r.Form["username"],r.Form["password"] )
                fmt.Fprintf(w,uStr)
		 }
}

func main()  {
	 http.HandleFunc("/login",login)
	 err := http.ListenAndServe(":9090",nil)
	 if err != nil{
		 log.Fatal("ListenAndServe",err)
	 }
}
