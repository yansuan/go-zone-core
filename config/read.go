package config

import "github.com/spf13/viper"

func Read(configFile string, inst interface{}) error {
	n := viper.New()
	n.SetConfigFile(configFile)
	err := n.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}
