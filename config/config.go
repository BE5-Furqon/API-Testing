package config

import (
	"sync"

	"github.com/labstack/gommon/log"

	"github.com/spf13/viper"
)

type AppConfig struct {
	Port string `yaml:"port"`
	Database struct {
		Driver string `yaml:"driver"`
		Name string `yaml:"name"`
		Host string `yaml:"host"`
		Port string `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	}
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

func GetConfig() *AppConfig {
	lock.Lock()
	defer lock.Unlock()

	if appConfig == nil {
		appConfig = initConfig()
	}

	return appConfig
}

func initConfig() *AppConfig {
	var defaultConfig AppConfig
	defaultConfig.Port = "1323"
	defaultConfig.Database.Driver = "mysql"
	defaultConfig.Database.Name = "crud_go"
	defaultConfig.Database.Host = "localhost"
	defaultConfig.Database.Port = "3306"
	defaultConfig.Database.Username = "root"
	defaultConfig.Database.Password = ""

	viper.SetConfigFile("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath("./config/")

	if err := viper.ReadInConfig(); err != nil {
		log.Info("Please make a config file!")
		return &defaultConfig
	}

	var finalConfig AppConfig

	err := viper.Unmarshal(&finalConfig)

	if err != nil {
		panic("Failed extract external config!")
	}

	return &finalConfig
}