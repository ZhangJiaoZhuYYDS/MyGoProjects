// @Author zhangjiaozhu 2023/3/17 11:14:00
package redis

import "github.com/go-redis/redis/v8"

var RDB = InitRedisDb()

func InitRedisDb() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
}
