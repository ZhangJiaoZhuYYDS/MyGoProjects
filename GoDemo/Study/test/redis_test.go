// @Author zhangjiaozhu 2023/3/17 10:55:00
package test

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"testing"
	"time"
)

var ctx = context.Background()

var rdb = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
	Password: "",
	DB: 0,
})

func TestRedisSet(t *testing.T)  {
	rdb.Set(ctx,"mykey","666",time.Second*10)
}

func TestRedisGet(t *testing.T)  {
	key,err := rdb.Get(ctx,"mykey").Result()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(key)
}