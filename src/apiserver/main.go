package main

import (
	"github.com/gin-gonic/gin"
	"apiserver/router"
	"net/http"
	"log"
	"time"
	"errors"
)

const (
	   BASEURL = "http://127.0.0.1"
	   PORT = ":8080"
)

func main()  {
	  gin.SetMode(gin.ReleaseMode)
      g := gin.New()

	middlewares :=[]gin.HandlerFunc{}

    router.Load(g,middlewares...)

	go func() {
		 if err := pingServer(); err != nil{
		 	 log.Fatal("The router has no response, or it might took too long to start up.",err)
		 }
	   	 log.Print("The router has been deployed successfully.")
	}()

	log.Printf("Start to listening the incoming requests on http address: %s", ":8080")
	log.Printf(http.ListenAndServe(PORT,g).Error())

 }


func pingServer() error {
	    for i := 0; i < 2;i++{
	    	resp,err := http.Get(BASEURL+PORT+"/sd/health")
	    	if err == nil && resp.StatusCode == 200{
	    		return nil
			}
			log.Print("Waiting for the router, retry in 1 second.")
	    	time.Sleep(time.Second)
		}
		return errors.New("Cannot connect to the router.")
}
