package connection

// import (
// 	"github.com/go-redis/redis"
// 	"github.com/spf13/viper"
// )

// var RedisClient *redis.Client

// func InitRedisClient() {
// 	RedisClient = redis.NewClient(&redis.Options{
// 		Addr:     viper.GetString("redis.host"),
// 		Password: viper.GetString("redis.password"),
// 		DB:       viper.GetInt("redis.db"),
// 	})
// }

// func CloseRedisClient() {
// 	_ = RedisClient.Close()
// }
