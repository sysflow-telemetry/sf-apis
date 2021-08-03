//
// Copyright (C) 2020 IBM Corporation.
//
// Authors:
// Frederico Araujo <frederico.araujo@ibm.com>
// Teryl Taylor <terylt@ibm.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package config implements configuration settings facilities.
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
