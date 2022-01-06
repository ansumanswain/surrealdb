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

import "time"

const (
	// Ignore specifies an ignored field
	Ignore = ignore
	// Prefix is the lowest char found in a key
	Prefix = prefix
	// Suffix is the highest char found in a key
	Suffix = suffix
	// Ignore specifies an ignored field
	ignore = "\x00"
	// Prefix is the lowest char found in a key
	prefix = "\x01"
	// Suffix is the highest char found in a key
	suffix = "\xff"
)

var (
	// StartOfTime is a datetime in the past
	StartOfTime = time.Unix(0, 0)
	// EndOfTime is a datetime in the future
	EndOfTime = time.Now().AddDate(50, 0, 0)
)

var (
	bEND = byte(0x00)
	bPRE = byte(0x01)
	bNIL = byte(0x02)
	bVAL = byte(0x03)
	bTME = byte(0x04)
	bNEG = byte(0x05)
	bPOS = byte(0x06)
	bSTR = byte(0x07)
	bARR = byte(0x08)
	bSUF = byte(0x09)
)

const (
	// MinNumber is the minimum number which can be accurately serialized
	MinNumber = -1 << 53
	// MaxNumber is the maximum number which can be accurately serialized
	MaxNumber = 1<<53 - 1
)

// Key ...
type Key interface {
	String() string
	Encode() []byte
	Decode(data []byte)
}