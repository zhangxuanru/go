package main

import "fmt"

func main()  {
	s := "hello world"
	newS := updateString(s)
	fmt.Println(newS)

	newS = joinStr("hello","Zxr")
	fmt.Println(newS)

	newS = updateStr("hello")
	fmt.Println(newS)

	//如果要声明一个多行的字符串,通过`来声明
	m := `hello
              world`
    fmt.Println(m)


}

/**
将字符串转为数组，并修改第一个字符
 */
func updateString(s string) string  {
    c := []byte(s)
    c[0] = 'Z'
    s2 := string(c)
    return  s2
}

//Go中可以使用+操作符来连接两个字符串：
func joinStr(str1,str2 string) string  {
    return  str1+str2
}

//修改字符串也可写为
func updateStr(str string) string  {
    str  = "c" + str[1:]  // 字符串虽不能更改，但可进行切片操作
    return str
}


