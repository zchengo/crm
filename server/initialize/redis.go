package initialize

import (
	"crm/global"
	"fmt"

	"github.com/go-redis/redis/v9"
)

// 初始化Redis数据库
func Redis() {
	r := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%v", r.Host, r.Port),
		Password: r.Password,
		DB:       r.Database,
	})
	global.Rdb = rdb
}
