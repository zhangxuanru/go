package main

import (
	"regexp"
	"fmt"
)

func main()  {
	isip := isIP("192.168.3.4")
	fmt.Println(isip)
	isInt := isInt("9")
	fmt.Println(isInt)
}


func isIP(ip string) bool  {
	  if m,_ := regexp.MatchString("^[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}$",ip);!m{
	  	 return  false
	  }
	  return true
}

func isInt(num string) bool {
	   m,_ := regexp.MatchString("^[0-9]+$",num)
       if !m{
             return false
	   }
	   return true
}
