package config

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

// Config map object that contains configuration attributes.
type Config map[string]string

// GetConfig returns a viper.Viper configuration object given a json file path.
func GetConfig(configFile string) (*viper.Viper, error) {
	s, err := os.Stat(configFile)
	if os.IsNotExist(err) {
		return nil, err
	}
	if s.IsDir() {
		return nil, errors.New("Config file is not a file")
	}
	c := viper.New()
	dir := filepath.Dir(configFile)
	c.SetConfigName(strings.TrimSuffix(filepath.Base(configFile), filepath.Ext(configFile)))
	c.SetConfigType("json")
	c.AddConfigPath(dir)
	err = c.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return c, nil
}
