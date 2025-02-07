package config

type Config struct {
	ZapConfig
	DatabaseConfig
	JwtConfig
}
type ZapConfig struct {
	Filename   string
	MaxSize    int
	MaxAge     int
	MaxBackups int
}
type DatabaseConfig struct {
	MysqlConfig
	RedisConfig
	MongoConfig
}
type MysqlConfig struct {
	Username string
	Password string
	Addr     string
	DB       string
}
type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}
type MongoConfig struct {
	Username string
	Password string
	Addr     string
	DB       string
}
type JwtConfig struct {
	JwtSecretKey string
	Issuer       string
}
