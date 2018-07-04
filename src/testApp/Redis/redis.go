package main

//redis 测试
import (
	"github.com/garyburd/redigo/redis"
	"time"
	"os"
	"os/signal"
	"syscall"
	"fmt"
)

var (
	pool *redis.Pool
)

func init()  {
	 redisHost := "127.0.0.1:6379"
	 pool = newPool(redisHost)
	 close()
}

func newPool(server string) *redis.Pool  {
	  return  &redis.Pool{
            MaxIdle:3,
            IdleTimeout:240*time.Second,

            Dial: func() (redis.Conn, error) {
				c,err := redis.Dial("tcp",server)
				if err != nil{
					return nil,err
				}
				return c, nil
			},
			TestOnBorrow: func(c redis.Conn, t time.Time) error {
				 _,err := c.Do("PING")
				 return  err
			},
	  }
}

func close()  {
     c := make(chan os.Signal,1)
	 signal.Notify(c,os.Interrupt)
     signal.Notify(c,syscall.SIGTERM)
     signal.Notify(c,syscall.SIGKILL)
     go func() {
     	<- c
     	pool.Close()
     	os.Exit(0)
	 }()
}

func Get(key string) ([]byte,error)  {
  conn := pool.Get()
  defer pool.Close()

  var data []byte
  data,err := redis.Bytes(conn.Do("Get",key))
  if err != nil{
		return  data, fmt.Errorf("error get key %s: %v", key, err)
   }
	return data, err
}

func main()  {
	test,error := Get("test")
	fmt.Println(test, error)
}

