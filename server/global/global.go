package global

import (
	"crm/config"

	"github.com/go-redis/redis/v9"
	"gorm.io/gorm"
)

var (
	Config config.Config
	Db     *gorm.DB
	Rdb    *redis.Client
)
