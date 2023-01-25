package configs

import (
	"github.com/spf13/viper"
)

// Configs root of application
type Configs struct {
	Server   Receiver    `mapstructure:"server"`
	Database Broker      `mapstructure:"database"`
	Assets   Destination `mapstructure:"assets"`
}

// Receiver configs
type Receiver struct {
	HTTPHost string `mapstructure:"http_host"`
	HTTPPort string `mapstructure:"http_port"`
}

// Broker configs
type Broker struct {
	HTTPHost string `mapstructure:"http_host"`
	HTTPPort string `mapstructure:"http_port"`
}

// Destination configs
type Destination struct {
	HTTPHost string `mapstructure:"http_host"`
	HTTPPort string `mapstructure:"http_port"`
}

// Init reads all configs from its yaml file
func Init(filePath string, defaultValues string) (*Configs, error) {
	v := viper.New()
	v.SetConfigType("yaml")

	v.SetConfigFile(filePath)
	v.AddConfigPath(".")
	v.AutomaticEnv()

	var cfg Configs
	if err := v.UnmarshalExact(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
