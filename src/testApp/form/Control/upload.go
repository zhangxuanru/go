package main

//上传文件处理
import (
	"net/http"
	"html/template"
	"fmt"
	"log"
	"time"
	"crypto/md5"
	"io"
	"strconv"
	"os"
)

const HTMLDIR = "src/testApp/form/tpl"
const UPLOADDIR = "src/testApp/form/upload/"

/*
上传文件
*/
func upload(w http.ResponseWriter, r *http.Request)  {
	if r.Method == "GET" {
        token := createToken()
		t,_ := template.ParseFiles(HTMLDIR+"/upload.gtpl")
		t.Execute(w,token)
	}
	//上传文件处理
	if r.Method == "POST" {
          r.ParseMultipartForm(32 << 20)
          file,handler,error := r.FormFile("uploadfile")
          if error != nil{
			  fmt.Println(error)
			  return
		  }
		  defer file.Close()

          fmt.Fprintf(w,"%v",handler.Header)

          f,err := os.OpenFile(UPLOADDIR+handler.Filename,os.O_WRONLY|os.O_CREATE,0666)
		 //f,err := os.Create(UPLOADDIR+handler.Filename)
          if err != nil{
			  fmt.Println("文件创建失败")
			  fmt.Println(err)
			  return
		  }
		  defer  f.Close()
		  io.Copy(f,file)
	}
}


/**
生成token
 */
func createToken() string {
	crutTime := time.Now().Unix()
	h := md5.New()
	io.WriteString(h,strconv.FormatInt(crutTime,10))
	token := fmt.Sprintf("%x",h.Sum(nil))  //生成md5token
	return  token
}


func main()  {
	http.HandleFunc("/upload",upload)
	err := http.ListenAndServe(":9090",nil)
	if err != nil{
		log.Fatal("ListenAndServe",err)
	}
}
