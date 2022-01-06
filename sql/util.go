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

package sql

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/hjson/hjson-go"
)

func in(token Token, tokens []Token) bool {

	for _, t := range tokens {
		if token == t {
			return true
		}
	}

	return false

}

func is(token Token, tokens ...Token) bool {

	for _, t := range tokens {
		if token == t {
			return true
		}
	}

	return false

}

func contains(search string, strings []string) bool {

	for _, str := range strings {
		if str == search {
			return true
		}
	}

	return false

}

func (p *parser) declare(tok Token, lit string) (interface{}, error) {

	if val := p.hold(tok); val != nil {
		return val, nil
	}

	switch tok {

	case TRUE:
		return true, nil

	case FALSE:
		return false, nil

	case MUL:
		return new(All), nil

	case QMARK:
		return new(Any), nil

	case NULL:
		return new(Null), nil

	case VOID:
		return new(Void), nil

	case MISSING:
		return new(Void), nil

	case EMPTY:
		return new(Empty), nil

	case EXPR:
		return NewIdent(lit), nil

	case IDENT:
		return NewIdent(lit), nil

	case TABLE:
		return NewTable(lit), nil

	case PARAM:
		return NewParam(lit), nil

	case REGEX:
		return NewRegex(lit), nil

	case STRING:
		return NewValue(lit), nil

	case REGION:
		return NewValue(lit), nil

	case DATE:
		return time.Parse(RFCDate, lit)

	case TIME:
		return time.Parse(RFCTime, lit)

	case NUMBER:
		val, err := strconv.ParseFloat(lit, 64)
		if err != nil {
			return val, fmt.Errorf("Invalid number: %s", lit)
		}
		return val, nil

	case DOUBLE:
		val, err := strconv.ParseFloat(lit, 64)
		if err != nil {
			return val, fmt.Errorf("Invalid number: %s", lit)
		}
		return val, nil

	case DURATION:
		var mul time.Duration
		switch {
		default:
			mul = 1
		case strings.HasSuffix(lit, "d"):
			mul, lit = 24, strings.Replace(lit, "d", "h", -1)
		case strings.HasSuffix(lit, "w"):
			mul, lit = 168, strings.Replace(lit, "w", "h", -1)
		}
		val, err := time.ParseDuration(lit)
		if err != nil {
			return val, fmt.Errorf("Invalid duration: %s", lit)
		}
		return val * mul, nil

	case ARRAY:
		var j []interface{}
		hjson.Unmarshal([]byte(lit), &j)
		if j == nil {
			return j, fmt.Errorf("Invalid JSON: %s", lit)
		}
		return j, nil

	case JSON:
		var j map[string]interface{}
		hjson.Unmarshal([]byte(lit), &j)
		if j == nil {
			return j, fmt.Errorf("Invalid JSON: %s", lit)
		}
		return j, nil

	}

	return lit, nil

}