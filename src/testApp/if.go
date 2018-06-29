package main

import "fmt"

func main()  {
	ifVarData()
	testFor()
	getMapData()
	switchData()
	switchFall()
}

func ifVarData()  {
	//条件判断语句里面允许声明一个变量，这个变量的作用域只能在该条件逻辑块内，其他地方就不起作用
	 if x := 9; x >10{
		 fmt.Println("x is greater than 10")
	 }else{
		 fmt.Println("x is less than 10")
	 }
}

//for
func testFor()  {
	  sum := 0
	  //for 循环，3个表达式都有的情况下
	  for index :=0;index < 10; index++{
	  	 sum += index
	  }
	  fmt.Println("sum is equal to ", sum)

	  //用while的方式求1-10的和
	  sum = 0
	  i := 0
	  for i < 10{
	  	 sum +=i
	  	  i++
	  }
	 fmt.Println("sum is while  equal to ", sum)


	  //忽略第一和第三表达式
	  sum = 1
	  for ; sum < 1000;{
			 sum += sum
		}
	  fmt.Println("sum is equal to to", sum)

	  //for 与 while  这种写法相当于while循环了
	  sum = 1
	  for sum < 1000{
	  	sum+=sum
	  }
	  fmt.Println("sum is equal to to to　", sum)

}

//for配合range可以用于读取slice和map的数据：
func getMapData()  {
        mapData := map[string]int{"a":10,"b":20,"c":30}
        //读取map
        for k,v:=range mapData{
			fmt.Println("map's key:",k)
			fmt.Println("map's val:",v)
		}
		//读取slice动态数组
		var arr []int
		arr = append(arr,100)
		arr = append(arr,90)
		arr = append(arr,80)
		for k,v := range arr{
			fmt.Println("arr's key:",k)
			fmt.Println("arr's val:",v)
		}

		//可以使用_来丢弃不需要的返回值 ,这里不获取key,只获取val
		for _, v := range mapData{
			fmt.Println("map's val:", v)
		}

}

/*
switch
Go里面switch默认相当于每个case最后带有break，匹配成功后不会自动向下执行其他case，而是跳出整个switch
*/
func switchData()  {
	i := 10
	switch i {
	case 1:
		fmt.Println("i is equal to 1")
	case 2,3,4: //可以很多值聚合在了一个case里面
		fmt.Println("i is equal to 2, 3 or 4")
	case 10:
		fmt.Println("i is equal to 10")
	default:
		fmt.Println("All I know is that i is an integer")
	}
}

/*
可以使用fallthrough强制执行后面的case代码

默认当case到一个就Break了， 加上 fallthrough，则还会执行下面的case 进行判断
*/
func switchFall()  {
	f := 6
	switch f {
	case 4:
		fmt.Println("The integer was <= 4")
	    fallthrough
	case 5:
		fmt.Println("The integer was <= 5")
	    fallthrough
	case 6:
		fmt.Println("The integer was <= 6")
		fallthrough
	case 7:
		fmt.Println("The integer was <= 7")
		fallthrough
	case 8:
		fmt.Println("The integer was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}
}
