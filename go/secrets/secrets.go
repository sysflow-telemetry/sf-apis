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

// Package secrets implements secret vault accessors.
package secrets

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/sysflow-telemetry/sf-apis/go/sfgo"
)

// Secrets stores a container secrets.
type Secrets struct {
	secretsDir string
	secrets    map[string]string
}

// NewSecrets creates an instance of Secrets with container secrets mounted to /run/secrets.
func NewSecrets() (*Secrets, error) {
	return NewSecretsWithCustomPath("/run/secrets")
}

// NewSecretsWithCustomPath creates an instance of Secrets with a custom secrets mount path.
func NewSecretsWithCustomPath(secretsDir string) (*Secrets, error) {
	if err := isDir(secretsDir); err != nil {
		return nil, err
	}
	return &Secrets{secretsDir: secretsDir, secrets: map[string]string{}}, nil
}

// Get reads secret value corresponding to key.
func (s *Secrets) Get(key string) (secret string, err error) {
	secret, err = s.read(key)
	if err != nil {
		return
	}
	return secret, nil
}

// GetDecoded reads and decode the base64 secret value corresponding to key.
func (s *Secrets) GetDecoded(key string) (string, error) {
	secret, err := s.read(key)
	if err != nil {
		return secret, err
	}
	decoded, err := base64.StdEncoding.DecodeString(secret)
	if err != nil {
		return sfgo.Zeros.String, err
	}
	return string(decoded), nil
}

// Reads a secret.
func (s *Secrets) read(secret string) (string, error) {
	if v, ok := s.secrets[secret]; ok {
		return v, nil
	}
	buf, err := ioutil.ReadFile(s.secretsDir + "/" + secret)
	if err != nil {
		return sfgo.Zeros.String, fmt.Errorf("secret %s does not exist or cannot be read: %v", secret, err)
	}
	v := strings.TrimSpace(string(buf))
	s.secrets[secret] = v
	return v, nil
}

// Checks if the given path is a directory. Returns nil if directory.
func isDir(path string) error {
	if fi, err := os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("Path %s not found", path)
	} else if !fi.Mode().IsDir() {
		return fmt.Errorf("Path %s is not a directory", path)
	}
	return nil
}
