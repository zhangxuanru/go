package main

import (
	"strings"
	"fmt"
)

func main()  {
	stringTest()
}

func stringTest()  {
	//字符串s中是否包含substr，返回bool值
	 s,substr := "abcdefg","b"
	isExists := strings.Contains(s,substr)
	fmt.Println("判断"+substr+"是否在"+s+"的结果为:",isExists)

	//字符串链接，把slice a通过sep链接起来
	strArr := []string{"hello","world","zxr"}
	fmt.Println(strings.Join(strArr,","))

	//在字符串s中查找sep所在的位置，返回位置值，找不到返回-1
	fmt.Println("x在zhangxuanrur 的索引位置为:",strings.Index("zhangxuanru","x"))


	//重复s字符串count次，最后返回重复的字符串
	fmt.Println(strings.Repeat("zxr",5))

	//在ss字符串中，把old字符串替换为new字符串，n表示替换的次数，小于0表示全部替换
	ss := "good good good zhang"
	fmt.Println(strings.Replace(ss,"zhang","zxr",-1))

	//把s字符串按照sep分割，返回slice
	fmt.Printf("%q%q\n",strings.Split("a,b,c",","),strings.Split(" xyz ", ""))

	//在s字符串的头部和尾部去除cutset指定的字符串
	fmt.Printf("[%q]",strings.Trim(" !!! Achtung !!! ","! "))

	//去除s字符串的空格符，并且按照空格分割返回slice
	fmt.Printf("Fields are: %q",strings.Fields("we are you doing ?"))

}