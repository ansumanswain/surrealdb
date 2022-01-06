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

func (p *parser) parseDefineStatement() (Statement, error) {

	// Inspect the next token.
	tok, _, err := p.shouldBe(NAMESPACE, DATABASE, LOGIN, TOKEN, SCOPE, TABLE, EVENT, FIELD, INDEX)

	switch tok {
	case NAMESPACE:
		return p.parseDefineNamespaceStatement()
	case DATABASE:
		return p.parseDefineDatabaseStatement()
	case LOGIN:
		return p.parseDefineLoginStatement()
	case TOKEN:
		return p.parseDefineTokenStatement()
	case SCOPE:
		return p.parseDefineScopeStatement()
	case TABLE:
		return p.parseDefineTableStatement()
	case EVENT:
		return p.parseDefineEventStatement()
	case FIELD:
		return p.parseDefineFieldStatement()
	case INDEX:
		return p.parseDefineIndexStatement()
	default:
		return nil, err
	}

}