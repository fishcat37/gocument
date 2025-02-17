package initialize

import (
	//"github.com/go-redis/redis/v8"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gocument/app/api/global"
	"gocument/app/api/internal/model"
	"golang.org/x/net/context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDatabase() {
	SetupMysql()
	SetupRedis()
	SetupMongo()
}

func SetupMysql() {
	mysqlConfig := global.Config.DatabaseConfig.MysqlConfig
	dsn := mysqlConfig.Username + ":" + mysqlConfig.Password + "@tcp(" + mysqlConfig.Addr + ")/" + mysqlConfig.DB + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		global.Logger.Fatal("failed to connect mysql")
	}
	global.MysqlDB = db
	err = db.AutoMigrate(&model.User{}, &model.Document{})
	if err != nil {
		global.Logger.Error("自动迁移表结构失败")
	}
	global.Logger.Info("mysql connect success")
}
func SetupRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     global.Config.RedisConfig.Addr,
		Password: global.Config.RedisConfig.Password,
		DB:       global.Config.RedisConfig.DB,
	})
	ctx := context.Background()
	_, err := rdb.Ping(ctx).Result()
	// defer func(rdb *redis.Client) {
	// 	if err := rdb.Close(); err != nil {
	// 		global.Logger.Error("failed to close redis connection")
	// 	}
	// }(rdb)
	if err != nil {
		global.Logger.Fatal("redis connect fail")
	}
	global.RedisDB = rdb
	global.Logger.Info("redis connect success")
}
func SetupMongo() {
	// TODO
	clientOptions := options.Client().ApplyURI("mongodb://" + global.Config.DatabaseConfig.MongoConfig.Username + ":" + global.Config.DatabaseConfig.MongoConfig.Password + "@" + global.Config.DatabaseConfig.MongoConfig.Addr)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	// defer func(client *mongo.Client) {
	// 	if err := client.Disconnect(context.TODO()); err != nil {
	// 		global.Logger.Error("failed to close mongodb connection")
	// 	}
	// }(client)
	if err != nil {
		global.Logger.Fatal("failed to connect mongodb")
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		global.Logger.Fatal("failed to ping mongodb")
	}
	global.MongoDB = client
	global.Logger.Info("mongodb connect success")
}
