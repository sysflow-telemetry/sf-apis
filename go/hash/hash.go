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
//
package hash

import (
	"fmt"

	"github.com/cespare/xxhash"
)

// GetHash computes the hash of its input arguments.
func GetHash(objs ...interface{}) uint64 {
	h := xxhash.New()
	for _, o := range objs {
		h.Write([]byte(fmt.Sprintf("%v", o)))
	}
	return h.Sum64()
}

// GetHashStr computes the hash string of its input arguments.
func GetHashStr(objs ...interface{}) string {
	h := xxhash.New()
	for _, o := range objs {
		h.Write([]byte(fmt.Sprintf("%v", o)))
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}
