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
//	Key string
	Total_size uint32
	Offset     uint32
	Tstamp     int32
	Ver int32
}

type KeyDir struct {
	map_ map[string]KeyEntry
	//entrys []KeyEntry
}

func NewKeyDir() (*KeyDir) {
	//var entrys []KeyEntry
	map_ := make(map[string]KeyEntry)
	return &KeyDir{map_}
}

func (dir *KeyDir) Set(key string, total_sz, offset uint32, tstamp, ver int32) error {
	if dir == nil {
		return ErrInvalid
	}

	entry := KeyEntry{total_sz, offset, tstamp, ver}
	dir.map_[key] = entry
	
	return nil
}

func (dir *KeyDir) Get(key string) (*KeyEntry, error) {
	if dir == nil {
		return nil, ErrInvalid
	}

	entry, ok := dir.map_[key]

	if ok {
		return &entry, nil		
	}

	return nil, errors.New("not exists.")
}

func (dir *KeyDir) Delete(key string) error {
	if dir == nil {
		return ErrInvalid
	}

	delete(dir.map_, key)

	return nil
}