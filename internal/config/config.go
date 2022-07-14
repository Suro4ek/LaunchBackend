package config

import (
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	GRPCHost string `env:"GRPC_HOST" yaml:"grpc_host"`
	GRPCPort string `env:"GRPC_PORT" yaml:"grpc_port"`
	Port     string `env:"PORT" yaml:"port"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{}
		if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
			cleanenv.GetDescription(instance, nil)
		}
	})
	return instance
}
