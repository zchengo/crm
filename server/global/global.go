package global

import (
	"crm/config"

	"github.com/go-redis/redis/v9"
	"github.com/smartwalle/alipay/v3"
	"gorm.io/gorm"
)

var (
	Config config.Config
	Db     *gorm.DB
	Rdb    *redis.Client
	Alipay *alipay.Client
)
