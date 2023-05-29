package user

import (
	"log"

	"github.com/spf13/viper"
)

type DBConfig struct {
	DBUsername string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
}

func GetDBConfig() *DBConfig {
	viper.SetConfigName("DBConfig")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("DB 설정 파일 읽기 실패: %v", err)
	}

	return &DBConfig{
		DBUsername: viper.GetString("db.username"),
		DBPassword: viper.GetString("db.password"),
		DBHost:     viper.GetString("db.host"),
		DBPort:     viper.GetString("db.port"),
		DBName:     viper.GetString("db.name"),
	}
}
