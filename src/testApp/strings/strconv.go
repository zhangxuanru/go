package main

import (
	"strconv"
	"fmt"
)

func main()  {
	strconvStr()
}

func strconvStr()  {

	//Append 系列函数将整数等转换为字符串后，添加到现有的字节数组中。
	str := make([]byte,0,10)
	str = strconv.AppendInt(str,4567, 10)
	str = strconv.AppendBool(str,false)
    str = strconv.AppendQuote(str,"abc")
    str = strconv.AppendQuoteRune(str,'单')
    fmt.Println(string(str))

      //Format 系列函数把其他类型的转换为字符串
       a := strconv.FormatBool(false)
       b := strconv.FormatFloat(123.23,'g',12,64)
       c := strconv.FormatInt(1234,10)
       d := strconv.FormatUint(1234,10)
       e := strconv.Itoa(1203)
	   fmt.Println(a, b, c, d, e)


       //Parse 系列函数把字符串转换为其他类型
       aa,_:= strconv.ParseBool("false")
       bb,_ := strconv.ParseFloat("123.64",64)
       cc,_ := strconv.ParseInt("1234",10,64)
       dd,_ := strconv.ParseUint("12345",10,64)
       ee,_ := strconv.Atoi("1023")
       fmt.Println(aa,bb,cc,dd,ee)

}
