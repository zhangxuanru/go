package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func main()  {
	//html := getContents()

//	fmt.Println(html)

	testPosix()
	
}

/*
抓取百度网页内容，并用正则去除一些格式的字符串
*/
func getContents()  string{
	   resp,err := http.Get("http://www.baidu.com")
	   if err != nil {
		   fmt.Println("http get error.")
	  }
	  defer resp.Body.Close()

	   body,err := ioutil.ReadAll(resp.Body)

		if err != nil {
			fmt.Println("http read error")
			return ""
		}

		src := string(body)

	   //将HTML标签全转换成小写
		re,_ := regexp.Compile("\\<[\\S\\s]+?\\>")
		src = re.ReplaceAllStringFunc(src,strings.ToLower)

	    //去除STYLE
	    re,_ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	    src = re.ReplaceAllString(src,"")

	    //去除SCRIPT
        re,_ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
        src = re.ReplaceAllString(src,"")

        //去除所有尖括号内的HTML代码，并换成换行符
        re,_ = regexp.Compile("\\<[\\S\\s]+?\\>")
        src = re.ReplaceAllString(src,"\n")

	    //去除连续的换行符
        re,_ = regexp.Compile("\\s{2,}")
        src = re.ReplaceAllString(src,"\n")

        return strings.TrimSpace(src)
}

func testPosix()  {
	 str := "aa09aaa88aaaa"
	 re,_ := regexp.Compile("[a-z]{2,4}")
	//查找符合正则的第一个
	one := re.Find([]byte(str))
	fmt.Println("Find:", string(one))

	all := re.FindAll([]byte(str),-1)
	fmt.Println("FindAll", all)

	//查找符合条件的index位置,开始位置和结束位置
	index := re.FindIndex([]byte(str))
	fmt.Println("FindIndex", index)

	//查找符合条件的所有的index位置，n同上
	allindex := re.FindAllIndex([]byte(str), -1)
	fmt.Println("FindAllIndex", allindex)


	a := "I am learning Go language"
	re2, _ := regexp.Compile("am(.*)lang(.*)")
	submatch := re2.FindSubmatch([]byte(a))
	fmt.Println("FindSubmatch", submatch)
	for _,v := range submatch{
		fmt.Println(string(v))
	}

	////定义和上面的FindIndex一样
	submatchindex := re2.FindSubmatchIndex([]byte(a))
	fmt.Println(submatchindex)

	//FindAllSubmatch,查找所有符合条件的子匹配
	submatchall := re2.FindAllSubmatch([]byte(a),-1)
	fmt.Println(submatchall)

	//FindAllSubmatchIndex,查找所有字匹配的index
	submatchallindex := re2.FindAllSubmatchIndex([]byte(a), -1)
	fmt.Println(submatchallindex)


}