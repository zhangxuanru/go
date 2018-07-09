package main

import (
	"html/template"
	"os"
	"fmt"
)

type Person struct {
	UserName string    //导出的字段，首字母是大写的
}

type Friend struct {
	   Fname string
}

type Product struct {
	UserName string
	Emails   []string
	Friends  []*Friend
}


func main()  {
	  ParseExecString()
	  fmt.Println("\n")
	  ParseThemeString()
}

/**
 解析模板字符串
 */
func ParseExecString(){
     t := template.New("fields name test")
     t, _ = t.Parse("hello {{.UserName}}")
     p := Person{UserName:"ZXR"}
     t.Execute(os.Stdout,p)
}

/*
模板中嵌套循环
*/
func ParseThemeString()  {
	  f1 := Friend{"hello"}
	  f2 := Friend{"world"}
	  t := template.New("test")
	  t,_ = t.Parse(`this {{.UserName}}
                      {{range .Emails}}
                           an Email {{.}}
                      {{end}}
                      {{with .Friends}}
                       {{range .}}
                            my friend name is {{.Fname}}
                      {{end}}
                      {{end}} 
                    `)
	  p := Product{UserName:"ZhangXuanRu",
		           Emails : []string{"1107099011@qq.com","strive965432@gmail.com"},
		           Friends: []*Friend{&f1,&f2},
	      }
   t.Execute(os.Stdout,p)
}