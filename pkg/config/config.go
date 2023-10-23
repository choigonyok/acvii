package config

import (
	"path/filepath"

	"github.com/choigonyok/goopt/pkg/env"
	"github.com/spf13/viper"
)

func init() {
	env.NewStringVar("ACVCTL", defaultConfig,
		"default values for acvii cli acvctl")
}

const (
	defaultConfig = "$HOME/.acvctl/config.yml"
)

var (
	Config = env.Get("CONFIG")
)

func InitConfigFromViper() error {
	configPath := filepath.Dir(Config.Name)
	baseName := filepath.Base(Config.Name)
	configType := filepath.Ext(Config.Name)
	configName := baseName[0 : len(baseName)-len(configType)]
	if configType != "" {
		configType = configType[1:]
	}
	viper.SetEnvPrefix("ACVCTL")
	viper.AutomaticEnv()
	viper.AllowEmptyEnv(true)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	viper.AddConfigPath(configPath)
	err := viper.ReadInConfig()
	if env.Get("ACVCTL").Name != defaultConfig {
		return err
	}
	return nil
}
