package main

import "fmt"

//声明一个结构体
type person1 struct {
	name string
	age int
}

func main()  {
	person()
	OlderAgeDiff()
	structTest()
	structCombination()
}

// 比较两个人的年龄，返回年龄大的那个人，并且返回年龄差
func Older(p1,p2 person1) (person1,int) {
	 if p1.age > p2.age{
	 	return  p1,p1.age-p2.age
	 }
	 return  p2,p2.age-p1.age
}

func OlderAgeDiff()  {
	 var tom person1
	 tom.name,tom.age = "hello",34
	 bob := person1{"join",18}
	// paus := person1{name:"paus",age:34}
	 tb_Older, tb_diff := Older(tom, bob)
	 fmt.Println(tb_Older,tb_diff)
}


//结构体练习
func person()  {
	//声明一个结构体
	type person struct {
		name string
		age int
	}

	//使用结构体 方法一
	var p person
	p.name = "张三"
	p.age = 10
	fmt.Printf("The person's name is %s\r\n",p.name)  // 访问P的name属性

	//使用结构体 方法二
	pp := person{"Tom",25}
	fmt.Printf("The person's name is %s age is %d\r\n",pp.name,pp.age)  // 访问Pp的name属性

	//使用结构体 方法三
	p1 := person{name:"李四",age:30}
	fmt.Printf("The person's name is %s age is %d\r\n",p1.name,p1.age)  // 访问p1的name属性

	//使用结构体 方法四
	p2 := new(person)
	p2.age = 80
	p2.name = "zxr"
	fmt.Printf("The person's name is %s age is %d\r\n",p2.name,p2.age)  // 访问p2的name属性



}




//struct的匿名字段
func structTest()  {
      //定义一个结构体
	type Human struct {
		  name string
		  age int
		  weight int
	  }

   //struct组合，Student组合了Human struct和string基本类型
   type Student struct {
   	    Human
	   speciality string
   }
    mark := Student{Human{"zxr",20,150},"hala"}
	// 我们访问相应的字段
	fmt.Println("His name is ", mark.name)
	fmt.Println("His age is ", mark.age)
	fmt.Println("His weight is ", mark.weight)
	fmt.Println("His speciality is ", mark.speciality)

}

/**
结构体 匿名访问， 组合的方式
 */
func structCombination()  {
	   type Skills []string
	   type Human struct {
		   name string
		   age int
		   weight int
	   }
	 type Student struct {
		Human   // 匿名字段，struct
		Skills  // 匿名字段，自定义的类型string slice
		int,   // 内置类型作为匿名字段
		speciality string
	}
	//// 初始化学生Jane
	jane := Student{Human:Human{"zhangxuanru",10,100},speciality:"haha"}
	// 现在我们来访问相应的字段
    fmt.Println(jane.name)
	jane.Skills = []string{"zhangsan"}
	jane.Skills = append(jane.Skills,"lisi","wangwu")
	fmt.Println("jane is skills = ",jane.Skills)
}
