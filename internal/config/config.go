package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"sync"
	"tAoD-advance/pkg/logging"
)

type Config struct {
	IsDebug *bool `yaml:"is_debug" env-default:"true"`
	Listen  struct {
		Type   string `yaml:"type" env-default:"port"`
		BindIp string `yaml:"bind_ip" env-default:"127.0.0.1"`
		Port   string `yaml:"port" env-default:"8080"`
	} `yaml:"listen"`
	MongoDB struct {
		Host       string `yaml:"host"  env-default:"localhost"`
		Port       string `yaml:"port"  env-default:"27017"`
		Database   string `yaml:"user_service"  env-default:"user_service"`
		Collection string `yaml:"collection"  env-default:"users"`
		Username   string `yaml:"username"`
		Password   string `yaml:"password"`
		AuthDB     string `yaml:"auth_db"`
	} `yaml:"mongo_db"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		logger := logging.GetLogger()
		logger.Info("read application configuration")
		instance = &Config{}

		if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Info(help)
			logger.Info(err)
		}
	})
	return instance
}
