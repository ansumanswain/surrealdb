// Copyright © 2016 SurrealDB Ltd.
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

package keys

// DB ...
type DB struct {
	KV string
	NS string
	DB string
}

// init initialises the key
func (k *DB) init() *DB {
	return k
}

// Copy creates a copy of the key
func (k *DB) Copy() *DB {
	return &DB{
		KV: k.KV,
		NS: k.NS,
		DB: k.DB,
	}
}

// Encode encodes the key into binary
func (k *DB) Encode() []byte {
	k.init()
	return encode(k.KV, "*", k.NS, "!", "d", k.DB)
}

// Decode decodes the key from binary
func (k *DB) Decode(data []byte) {
	k.init()
	var __ string
	decode(data, &k.KV, &__, &k.NS, &__, &__, &k.DB)
}

// String returns a string representation of the key
func (k *DB) String() string {
	k.init()
	return output(k.KV, "*", k.NS, "!", "d", k.DB)
}