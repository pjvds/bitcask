// bitcask: Eric Brewer-inspired key/value store, in Golang
//
// Copyright (c) 2014 Ji-Guorui<jiguorui@gmail.com>. All Rights Reserved.
//
// This file is provided to you under the Apache License,
// Version 2.0 (the "License"); you may not use this file
// except in compliance with the License.  You may obtain
// a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package bitcask

import (
	"errors"
)

type KeyEntry struct {
	Key string
	Total_size uint32
	Offset     uint32
	Tstamp     int32
	Ver int32
}

type KeyDir struct {
	entrys []KeyEntry
}

func NewKeyDir() (*KeyDir) {
	var entrys []KeyEntry
	return &KeyDir{entrys}
}

func (dir *KeyDir) Set(key string, total_sz, offset uint32, tstamp, ver int32) error {
	if dir == nil {
		return ErrInvalid
	}

	entry := KeyEntry{key, total_sz, offset, tstamp, ver}
	dir.entrys = append(dir.entrys, entry)
	return nil
}

func (dir *KeyDir) Get(key string) (*KeyEntry, error) {
	if dir == nil {
		return nil, ErrInvalid
	}

	sz := len(dir.entrys)
	for i := 0; i < sz; i++ {
		if key == dir.entrys[i].Key {
			return &(dir.entrys[i]), nil
		}
	}
	return nil, errors.New("not exists.")
}