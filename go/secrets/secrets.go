package secrets

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
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
	secrets := &Secrets{secretsDir: secretsDir, secrets: map[string]string{}}
	err := secrets.readAll()
	return secrets, err
}

// Get reads secret value corresponding to key.
func (s *Secrets) Get(key string) (string, error) {
	if _, ok := s.secrets[key]; !ok {
		return "", fmt.Errorf("secret %s does not exist", key)
	}
	return s.secrets[key], nil
}

// GetDecoded reads and decode the base64 secret value corresponding to key.
func (s *Secrets) GetDecoded(key string) ([]byte, error) {
	if _, ok := s.secrets[key]; !ok {
		return nil, fmt.Errorf("secret %s does not exist", key)
	}
	return base64.StdEncoding.DecodeString(s.secrets[key])
}

// GetAll returns all container secrets.
func (s *Secrets) GetAll() map[string]string {
	return s.secrets
}

// Reads all secrets from the secrets mount path.
func (s *Secrets) readAll() error {
	err := isDir(s.secretsDir)
	if err != nil {
		return err
	}
	secrets, err := ioutil.ReadDir(s.secretsDir)
	if err != nil {
		return err
	}
	for _, secret := range secrets {
		err := s.read(secret.Name())
		if err != nil {
			return err
		}
	}
	return nil
}

// Reads a secret.
func (s *Secrets) read(secret string) error {
	buf, err := ioutil.ReadFile(s.secretsDir + "/" + secret)
	if err != nil {
		return err
	}
	s.secrets[secret] = strings.TrimSpace(string(buf))
	return nil
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
