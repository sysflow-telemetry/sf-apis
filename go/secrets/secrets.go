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
func (s *Secrets) GetDecoded(key string) ([]byte, error) {
	secret, err := s.read(key)
	if err != nil {
		return nil, err
	}
	return base64.StdEncoding.DecodeString(secret)
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
