package config

import (
	"fmt"
	"strings"

	"github.com/google/wire"
	"github.com/spf13/viper"
)

// Init 初始化viper
func New(appName string) (*viper.Viper, error) {
	var (
		err error
		v   = viper.New()
	)
	v.SetEnvPrefix(strings.ToUpper(appName))
	v.AutomaticEnv()

	v.SetConfigName("config")              // name of config file (without extension)
	v.AddConfigPath("/etc/" + appName)  // path to look for the config file in
	v.AddConfigPath("$HOME/." + appName) // call multiple times to add many search paths
	v.AddConfigPath(".")                   // optionally look for config in the working directory

	if err := v.ReadInConfig(); err == nil {  // Find and read the config file
		fmt.Printf("use config file -> %s\n", v.ConfigFileUsed())
	} else {
		return nil, err
	}
	// global defaults
	v.SetDefault("json_logs", false)
	v.SetDefault("loglevel", "debug")

	return v, err
}

var ProviderSet = wire.NewSet(New)
