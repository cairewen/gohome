package redis

import (
	"github.com/go-redis/redis"
	"time"
	"log"
	"gohome/config"
)

var client *redis.Client

func init() {
	host := config.GetYamlValue("redis.host")
	client = redis.NewClient(&redis.Options{
		Addr:     host,
		Password: "",
		DB:       0,
	})
}

func Ping() {

}

/**
	put key value into redis forever
 */
func Put(key string, value interface{}) {
	PutX(key, value, 0)
}

/**
	put key value into redis witho expiration time
 */
func PutX(key string, value interface{}, expiration time.Duration) {
	err := client.Set(key, value, expiration)
	if err != nil {
		log.Println(err)
	}
}

/**
	get value with input key
 */
func Get(key string) interface{} {
	result, err := client.Get(key).Result()
	if err != nil {
		log.Println(err)
	}
	return result
}
