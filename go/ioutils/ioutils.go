//
// Copyright (C) 2021 IBM Corporation.
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

// Package ioutils implements IO utilities.
package ioutils

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// ListFilePaths lists file paths with extension fileExt in path if
// path is a valid directory, otherwise, it returns path if path is
// a valid path and has extension fileExt.
func ListFilePaths(path string, fileExt string) ([]string, error) {
	var paths []string
	if fi, err := os.Stat(path); os.IsNotExist(err) {
		return paths, err
	} else if fi.IsDir() {
		var files []os.FileInfo
		var err error
		if files, err = ioutil.ReadDir(path); err != nil {
			return paths, err
		}
		for _, file := range files {
			if filepath.Ext(file.Name()) == fileExt {
				f := path + "/" + file.Name()
				paths = append(paths, f)
			}
		}
		return paths, nil
	} else {
		if filepath.Ext(path) == fileExt {
			return append(paths, path), nil
		}
		return paths, nil
	}
}

//FileExists checks whether a file exists and whether it is a directory.
func FileExists(filename string) (bool, bool) {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false, false
	}
	return true, info.IsDir()
}
