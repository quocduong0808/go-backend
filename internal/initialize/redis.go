package initialize

import (
	"context"
	"fmt"
	"go/go-backend-api/global"

	"github.com/redis/go-redis/v9"
)

func InitRegis() {
	config := global.Config.Redis
	global.Redis = redis.NewClient(&redis.Options{
		Addr:           fmt.Sprintf("%s:%v", config.Host, config.Port),
		Password:       config.Password, // No password set
		DB:             config.Db,       // Use default DB
		PoolSize:       config.PoolSize,
		MaxActiveConns: config.MaxActiveConns,
		MaxIdleConns:   config.MaxIdleConns,
	})

	ctx := context.Background()

	err := global.Redis.Ping(ctx).Err()
	global.HandleErrorPanic(err, "error when create connection to redis")
	global.Logger.Info("create connection to redis success")

}
