package config

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func InitSingleRedis() {
	var redisConf Redis
	rConf := GlobalConfig.Sub("redis")
	redisMap := rConf.AllSettings()
	mapstructure.Decode(redisMap, &redisConf)
	redisAddr := fmt.Sprintf("%s:%d", redisConf.Host, redisConf.Port)
	slog.Info("连接数据：%s", redisAddr)
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisConf.Password,
		DB:       redisConf.Db,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	pong, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		slog.Info("redis connect failed, err:%v", err)
	} else {
		slog.Info("redis connect success, res:%s", pong)
	}

}
