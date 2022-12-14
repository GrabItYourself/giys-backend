package config

import (
	"strings"
	"sync"

	"github.com/GrabItYourself/giys-backend/lib/logger"
	"github.com/GrabItYourself/giys-backend/lib/postgres"
	"github.com/GrabItYourself/giys-backend/lib/rabbitmq"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var configOnce sync.Once
var config *Config

type Config struct {
	Server   ServerConfig    `mapstructure:"server"`
	Log      logger.Config   `mapstructure:"log"`
	Postgres postgres.Config `mapstructure:"postgres"`
	RabbitMQ rabbitmq.Config `mapstructure:"rabbitmq"`
	Grpc     GrpcConfig      `mapstructure:"grpc"`
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
}

type GrpcConfig struct {
	Payment GrpcConnection `mapstructure:"payment"`
}

type GrpcConnection struct {
	Addr string `mapstructure:"addr"`
}

func InitConfig() *Config {
	configOnce.Do(func() {

		viper.SetConfigName("config")                  // name of config file without extension
		viper.AddConfigPath("./order/internal/config") // path to look for config file, relative to working directory
		viper.AddConfigPath("/config")                 // production config mount path

		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		if err := viper.ReadInConfig(); err != nil {
			panic(errors.Wrap(err, "Config file not found"))
		}
		viper.AutomaticEnv()

		viper.WatchConfig() // Watch for changes to the configuration file and recompile
		if err := viper.Unmarshal(&config); err != nil {
			panic(errors.Wrap(err, "can't unmarshal config"))
		}
		logger.Info("Config initialized!")
	})
	return config
}
