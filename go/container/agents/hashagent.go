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

// Package agents implements an agent for introspecting in container filesystems. 
package agents

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io"
	"os"

	"github.com/spf13/viper"
	"github.com/sysflow-telemetry/sf-apis/go/container/storage"
	"github.com/sysflow-telemetry/sf-apis/go/ioutils"
)

// HashAgent represents the file hasher object for hashing files in running containers.
type HashAgent struct {
	store *storage.CStorage
}

// Constants
const (
	BUFFER_SIZE = 1024
)

// NewHashAgent constructs a new hash agent. given a configuration file.
func NewHashAgent(conf *viper.Viper) (*HashAgent, error) {
	cStore, err := storage.NewContainerStore(conf)
	if err != nil {
		return nil, err
	}
	return &HashAgent{cStore}, nil
}

// GetHashes  Calculates the md5, sha1, and sha256 hashes for the given filepath, and container ID.  Also returns the size of the
// file hashed, and the container layer on which the file was found:  host mount, container, or image layer.
func (h *HashAgent) GetHashes(containerID string, filePath string) (string, string, string, int, storage.Layer, error) {
	var s1s string
	var s256s string
	var m5s string
	path, layer, err := h.store.GetFilePath(containerID, filePath)
	if err != nil {
		return m5s, s1s, s256s, 0, layer, err
	} else if layer == storage.LUNKNOWN {
		return m5s, s1s, s256s, 0, layer, errors.New("FileNotFound")
	}

	if exists, dir := ioutils.FileExists(path); exists && !dir {
		buffer := make([]byte, BUFFER_SIZE)
		file, err := os.Open(path)
		if err != nil {
			return m5s, s1s, s256s, 0, layer, err
		}
		defer file.Close()
		totalBytes := 0
		m5 := md5.New()
		s1 := sha1.New()
		s256 := sha256.New()
		for {
			bytesread, err := file.Read(buffer)
			if err != nil {
				if err != io.EOF {
					return m5s, s1s, s256s, 0, layer, err
				}
				break
			}
			totalBytes += bytesread
			m5.Write(buffer[:bytesread])
			s1.Write(buffer[:bytesread])
			s256.Write(buffer[:bytesread])
		}
		return hex.EncodeToString(m5.Sum(nil)), hex.EncodeToString(s1.Sum(nil)), hex.EncodeToString(s256.Sum(nil)), totalBytes, layer, nil
	} else if dir {
		return m5s, s1s, s256s, 0, layer, errors.New("FileIsDir")
	}
	return m5s, s1s, s256s, 0, layer, errors.New("FileNotFound")
}
