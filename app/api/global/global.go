package global

import (
	"github.com/redis/go-redis/v9"
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
	//Clients map[*websocket.Conn]bool
	//Lock sync.Mutex
	//Upgrader   websocket.Upgrader
)
