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

// Package async implements asynchronous functions to compute file hashes in containers. 
package async

import (
	"path/filepath"
	"sync"
	"time"

	cache "github.com/patrickmn/go-cache"
	"github.com/spf13/viper"
	"github.com/sysflow-telemetry/sf-apis/go/container/agents"
	"github.com/sysflow-telemetry/sf-apis/go/container/storage"
	"github.com/sysflow-telemetry/sf-apis/go/logger"
)

// ErrorStatus represents the occurrence of an error during file hashing
type ErrorStatus uint8

// HashCallback is a callback method for when the hash is ready.
type HashCallback func(fhi *FileHashInfo)

// Constants
const (
	SUCCESS              ErrorStatus = 0
	FILENOTFOUND         ErrorStatus = 1
	FILEISDIR            ErrorStatus = 2
	CUSTOM               ErrorStatus = 3
	CACHE_EXPIRE                     = "cacheExpire"
	CACHE_PURGE                      = "cachePurge"
	NUM_HASH_THREADS                 = "numHashThreads"
	DEFAULT_EXP                      = 5
	DEFAULT_PUR                      = 7
	DEFAULT_HASH_THREADS             = 5
)

// FilePath  contains the path and container ID for the file to be hashed.
type FilePath struct {
	Path        string
	ContainerID string
}

// FileHashTable is a thread safe, asynchronous hash table with which manages file hashes.
type FileHashTable struct {
	hashTable  *cache.Cache
	hashChan   chan *FilePath
	hashAgent  *agents.HashAgent
	hashCb     HashCallback
	wg         *sync.WaitGroup
	numThreads int
}

// FileHashInfo stores the hashes, and information about the hashed file.
type FileHashInfo struct {
	File       *FilePath
	Status     ErrorStatus
	Error      error
	Md5        string
	Sha1       string
	Sha256     string
	Layer      storage.Layer
	LastUpdate time.Time
	Size       int
}

// NewFileHashTable creates a new asysnchronous file hash table given a configurtion and optional wait group.
func NewFileHashTable(conf *viper.Viper, hcb HashCallback) (*FileHashTable, error) {
	h := new(FileHashTable)
	var err error
	expire := DEFAULT_EXP
	purge := DEFAULT_PUR
	numThreads := DEFAULT_HASH_THREADS

	if conf.IsSet(CACHE_EXPIRE) {
		expire = conf.GetInt(CACHE_EXPIRE)
	}

	if conf.IsSet(CACHE_PURGE) {
		purge = conf.GetInt(CACHE_PURGE)
	}

	if conf.IsSet(NUM_HASH_THREADS) {
		numThreads = conf.GetInt(NUM_HASH_THREADS)
	}
	h.hashTable = cache.New(time.Duration(expire)*time.Minute, time.Duration(purge)*time.Minute)
	h.hashChan = make(chan *FilePath, numThreads)
	h.hashCb = hcb
	h.numThreads = numThreads
	hashAgent, err := agents.NewHashAgent(conf)
	h.wg = new(sync.WaitGroup)
	h.wg.Add(numThreads)

	if err != nil {
		return nil, err
	}
	h.hashAgent = hashAgent
	return h, nil
}

// Init initializes the file hashing thread.
func (f *FileHashTable) Init() {
	for i := 0; i < f.numThreads; i++ {
		go f.hashThread()
	}
}

// StartHash Asynchronous call to hash a file in a container.
func (f *FileHashTable) StartHash(fp *FilePath) {
	f.hashChan <- fp
}

// Close closes down the hashing thread.
func (f *FileHashTable) Close() {
	close(f.hashChan)
	f.wg.Wait()
}

// Get returns the file hashes of file in a container, if it exists.  Note that the file hash will
// only appear in the hash when the asynchronous thread has completed the file hashing calculations.
func (f *FileHashTable) Get(containerID string, pth string) *FileHashInfo {
	if entry, ok := f.hashTable.Get(containerID + filepath.Clean(pth)); ok {
		return entry.(*FileHashInfo)
	}
	return nil
}

// Remove removes an item from the cache if present.
func (f *FileHashTable) Remove(containerID string, pth string) {
	f.hashTable.Delete(containerID + filepath.Clean(pth))
}

// hashThread launches the hashing thread and is typically called as a go function.
func (f *FileHashTable) hashThread() {
	if f.wg != nil {
		defer f.wg.Done()
	}
	for {
		fp, ok := <-f.hashChan
		if !ok {
			logger.Trace.Println("Channel closed. Shutting down.")
			break
		}
		fhi := new(FileHashInfo)
		key := fp.ContainerID + filepath.Clean(fp.Path)
		md5, sha1, s256, bytes, layer, err := f.hashAgent.GetHashes(fp.ContainerID, fp.Path)
		fhi.File = fp
		if err != nil {
			logger.Error.Println("Error hashing file: " + fp.Path + " Error: " + err.Error())
			if err.Error() == "FileNotFound" {
				fhi.Status = FILENOTFOUND
				fhi.Error = err
			} else if err.Error() == "FileIsDir" {
				fhi.Status = FILEISDIR
				fhi.Error = err
			} else {
				fhi.Status = CUSTOM
				fhi.Error = err
			}
		} else {
			fhi.Status = SUCCESS
			fhi.Md5 = md5
			fhi.Sha1 = sha1
			fhi.Sha256 = s256
			fhi.Layer = layer
			fhi.Size = bytes
		}
		fhi.LastUpdate = time.Now()
		f.hashTable.Set(key, fhi, cache.DefaultExpiration)
		if f.hashCb != nil {
			f.hashCb(fhi)
		}
	}
}
