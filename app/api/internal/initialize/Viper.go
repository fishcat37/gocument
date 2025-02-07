package initialize

import (
	"github.com/spf13/viper"
	"gocument/app/api/global"
)

func SetupViper() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.SetConfigFile("./manifest/config.yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic("read config failed" + err.Error())
	}
	err = viper.Unmarshal(&global.Config)
	if err != nil {
		panic("unmarshal config failed" + err.Error())
	}
}
