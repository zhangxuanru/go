package main

import (
	"strconv"
	"fmt"
	"regexp"
	"html/template"
	"time"
)

/**
格式判断
 */
func validata()  {
	age := "23"
	//如果我们是判断正整数，那么我们先转化成int类型
	getInt,error := strconv.Atoi(age)
	if error != nil{
		fmt.Println("数字转化出错了 就不是数字")
	}

	//正则表达式判断整型
	 m,_ := regexp.MatchString("^[0-9]+$",age)
	 fmt.Println(getInt,m)
	 //电子邮件地址
	 //regexp.MatchString("^([\w\.\_]{2,10})@(\w{1,}).([a-zA-z]{2,4})$")

	 //手机号码
	 //regexp.MatchString("^(1[3|4|5|8][0-9]\d{4,8})$")

	 //转义字符串
	 str := template.HTMLEscapeString("<script>alert('erwerewr')</script>")
	 fmt.Println(str)

	 fmt.Println(time.Now().Unix())

	 //t,err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
	 //err = t.ExecuteTemplate(out, "T","<script>alert('you have been pwned')</script>")
	 //fmt.Println(err)

}


func main()  {
  validata()
}
