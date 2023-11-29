package redis

import (
	"go-edu/config"
	"go.uber.org/zap"
	"gopkg.in/redis.v5"
	"log"
	"time"
)

var redisCache *redis.Client

// 创建 redis 客户端
func createClient(redisHost string, password string, dataBase int, PoolSize int) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: password,
		DB:       dataBase,
		PoolSize: PoolSize,
	})

	// 通过 cient.Ping() 来检查是否成功连接到了 redis 服务器
	_, err := client.Ping().Result()
	if err != nil {
		zap.L().Error("连接失败" + err.Error())
	}

	return client
}

func Init(cfg *config.RedisConfig) (err error) {
	redisHost := cfg.Host
	dataBase := cfg.DB
	password := cfg.Password
	poolsize := cfg.PoolSize
	redisCache = createClient(redisHost, password, dataBase, poolsize)
	return
}

func SetStr(key, value string, time time.Duration) (err error) {
	err = redisCache.Set(key, value, time).Err()
	if err != nil {
		log.Print("set key:", key, ",value:", value, err)
	}
	return
}

func GetStr(key string) (value string) {
	v, _ := redisCache.Get(key).Result()
	return v
}

func DelKey(key string) (err error) {
	err = redisCache.Del(key).Err()
	return
}
