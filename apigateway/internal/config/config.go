package config

import (
	"strings"
	"sync"

	"github.com/GrabItYourself/giys-backend/lib/logger"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var configOnce sync.Once
var config *Config

type GrpcConnection struct {
	Addr string `mapstructure:"addr"`
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
}

type GRPCConfig struct {
	User  GrpcConnection `mapstructure:"user"`
	Auth  GrpcConnection `mapstructure:"auth"`
	Order GrpcConnection `mapstructure:"order"`
	Shop  GrpcConnection `mapstructure:"shop"`
}

type OAuthConfig struct {
	ClientId     string `mapstructure:"client_id"`
	ClientSecret string `mapstructure:"client_secret"`
	RedirectURL  string `mapstructure:"redirect_url"`
}

type Config struct {
	Server ServerConfig  `mapstructure:"server"`
	Log    logger.Config `mapstructure:"log"`
	Grpc   GRPCConfig    `mapstructure:"grpc"`
	OAuth  OAuthConfig   `mapstructure:"oauth"`
}

func InitConfig() *Config {
	configOnce.Do(func() {
		viper.SetConfigName("config")                       // name of config file without extension
		viper.AddConfigPath("./apigateway/internal/config") // path to look for config file, relative to working directory
		viper.AddConfigPath("/config")                      // production config mount path

		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		if err := viper.ReadInConfig(); err != nil {
			logger.Panic(errors.Wrap(err, "Config file not found").Error())
		}
		viper.AutomaticEnv()

		viper.WatchConfig() // Watch for changes to the configuration file and recompile
		if err := viper.Unmarshal(&config); err != nil {
			logger.Panic(errors.Wrap(err, "can't unmarshal config").Error())
		}
		logger.Info("Config initialized!")
	})
	return config
}
