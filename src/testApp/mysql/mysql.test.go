package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
)

func main()  {
	Insert()

	getSelectData()

	updateData()

	delData()

}



/**
  插入数据
 */
func Insert()  {
     db,err := sql.Open("mysql","root:abc123123@tcp(23.106.155.177:3306)/test?charset=utf8")
     checkErr(err)
     stmp,err := db.Prepare("INSERT INTO userinfo(username,department,created) VALUES (?,?,?)")
     checkErr(err)
     res,err := stmp.Exec("zxr","PHP","2018-07-02")
     checkErr(err)
     id,err := res.LastInsertId()
     checkErr(err)
     fmt.Println(id)
}

/**
修改语句
 */
func updateData()  {
    db,err := sql.Open("mysql","root:abc123123@tcp(23.106.155.177:3306)/test?charset=utf8")
    checkErr(err)
    stmp,err := db.Prepare("UPDATE userinfo SET username= ? where uid=?")
    checkErr(err)
    res,err := stmp.Exec("jiaojiao",9)
    checkErr(err)
    fmt.Println(res.RowsAffected())
}


/**
查询语句
 */
func getSelectData()  {
   db,err := sql.Open("mysql","root:abc123123@tcp(23.106.155.177:3306)/test?charset=utf8")
   checkErr(err)

   //方法一：查询预处理
   stmp,err := db.Prepare("SELECT * FROM userinfo WHERE uid =? ")
   rows,err :=  stmp.Query(9)

   //方法二，直接执行SQL语句
  // rows,err := db.Query("select * from userinfo where uid=1;")

   checkErr(err)
   for rows.Next(){
   	    var uid int
   	    var username string
   	    var department string
   	    var created string
   	    err = rows.Scan(&uid,&username,&department,&created)
   	    checkErr(err)
   	    fmt.Println(uid,username,department,created)
    }
}

/**
删除语句
 */
func delData()  {
    db,err := sql.Open("mysql","root:abc123123@tcp(23.106.155.177:3306)/test?charset=utf8")
    checkErr(err)
    stmp,err := db.Prepare("DELETE FROM userinfo WHERE uid=?")
    checkErr(err)
    res,err := stmp.Exec(1)
    checkErr(err)
    fmt.Println(res.RowsAffected())
}









func checkErr(err error)  {
	  if err != nil{
	   	 panic(err)
	  }
}
