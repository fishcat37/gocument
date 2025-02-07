package global

import (
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"gocument/app/api/global/config"
	"gorm.io/gorm"
)

var (
	Config  *config.Config
	Logger  *zap.Logger
	MysqlDB *gorm.DB
	RedisDB *redis.Client
	MongoDB *mongo.Client
)
