package main

import (
	"os"
	"fmt"
	"net"
)

func main()  {
	getArgs()
}

//ParseIP(s string) IP函数会把一个IPv4或者IPv6的地址转化成IP类型
func getArgs()  {
	if len(os.Args) != 2{
		fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]
	addr := net.ParseIP(name)

	if addr == nil {
		fmt.Println("Invalid address")
	} else {
		fmt.Println("The address is ", addr.String())
	}
	os.Exit(0)
}




