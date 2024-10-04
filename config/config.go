package config

import (
	"strings"
	"sync"
	
	"github.com/spf13/viper"
)

type(
	Config struct {
		// Configurations here
		Server *Server
		Database *Database
	}

	Server struct {
		// Server configurations here
		Port int
	}

	Database struct {
		// Database configurations here
		Host string
		Port int
		User string
		Password string
		DatabaseName string
		SslMode string
		Timezone string
	}
)

var (
	once sync.Once
	configInstance *config
)

func GetConfig() *config {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./")
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		
		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}

		if err := viper.Unmarshal(&configInstance); err != nil {
			panic(err)
		}
	})
	return configInstance
}