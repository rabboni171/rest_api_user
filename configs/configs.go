package configs

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Dbname   string `mapstructure:"dbname"`
	Password string `mapstructure:"password"`
	SSLMode  string `mapstructure:"sslmode"`
}

func InitConfig() (config Config, err error) {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(config)
	return config, nil
}
