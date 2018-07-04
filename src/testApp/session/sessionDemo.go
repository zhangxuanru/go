package main

import (
	"github.com/astaxie/beego/session"
	"net/http"
	"fmt"
)

var globalSessions *session.Manager

func init()  {
	  var config = &session.ManagerConfig{
		  CookieName:"gosessionid",
		  EnableSetCookie:true,
		  Gclifetime:3600,
		  ProviderConfig: "./tmp",
	  }
  	  globalSessions,_ = session.NewManager("file", config)
      go globalSessions.GC()
	}


func login(w http.ResponseWriter, r *http.Request)  {
	 sess,_:= globalSessions.SessionStart(w,r)
	 defer sess.SessionRelease(w)
	 username := sess.Get("username")
	 if username == nil{
	 	sess.Set("username","zhangxuanru")
	 }
	 fmt.Println(username)
	 fmt.Fprintf(w,"hello world 123,username:",username)
}



func main()  {
	 http.HandleFunc("/login", login)
	 err := http.ListenAndServe(":9100",nil)
	 if err != nil{
	 	fmt.Println(err)
	 }
}



