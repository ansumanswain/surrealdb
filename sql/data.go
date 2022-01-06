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

func (p *parser) parseData() (exp Expr, err error) {

	if tok, _, exi := p.mightBe(SET, DIFF, MERGE, CONTENT); exi {

		if is(tok, SET) {
			if exp, err = p.parseSet(); err != nil {
				return nil, err
			}
		}

		if is(tok, DIFF) {
			if exp, err = p.parseDiff(); err != nil {
				return nil, err
			}
		}

		if is(tok, MERGE) {
			if exp, err = p.parseMerge(); err != nil {
				return nil, err
			}
		}

		if is(tok, CONTENT) {
			if exp, err = p.parseContent(); err != nil {
				return nil, err
			}
		}

	}

	return

}

func (p *parser) parseSet() (mul Expr, err error) {

	out := &DataExpression{}

	for {

		var tok Token
		var lit string

		one := &ItemExpression{}

		// The first part of a SET expression must
		// always be an identifier, specifying a
		// record field to set.

		tok, lit, err = p.shouldBe(IDENT, EXPR)
		if err != nil {
			return nil, &ParseError{Found: lit, Expected: []string{"field name"}}
		}

		one.LHS, err = p.declare(tok, lit)
		if err != nil {
			return nil, err
		}

		// The next query part must be a =, +=, or
		// -= operator, as this is a SET expression
		// and not a binary expression.

		one.Op, lit, err = p.shouldBe(EQ, INC, DEC)
		if err != nil {
			return nil, err
		}

		// The next query part can be any expression
		// including a parenthesised expression or a
		// binary expression so handle accordingly.

		one.RHS, err = p.parseExpr()
		if err != nil {
			return nil, err
		}

		// Append the single SET data expression to
		// the array of data expressions.

		out.Data = append(out.Data, one)

		// Check to see if the next token is a comma
		// and if not, then break out of the loop,
		// otherwise repeat until we find no comma.

		if _, _, exi := p.mightBe(COMMA); !exi {
			break
		}

	}

	return out, err

}

func (p *parser) parseDiff() (exp *DiffExpression, err error) {

	exp = &DiffExpression{}

	tok, lit, err := p.shouldBe(ARRAY, PARAM)
	if err != nil {
		return nil, &ParseError{Found: lit, Expected: []string{"json"}}
	}

	exp.Data, err = p.declare(tok, lit)
	if err != nil {
		return nil, &ParseError{Found: lit, Expected: []string{"json"}}
	}

	return

}

func (p *parser) parseMerge() (exp *MergeExpression, err error) {

	exp = &MergeExpression{}

	tok, lit, err := p.shouldBe(JSON, PARAM)
	if err != nil {
		return nil, &ParseError{Found: lit, Expected: []string{"json"}}
	}

	exp.Data, err = p.declare(tok, lit)
	if err != nil {
		return nil, &ParseError{Found: lit, Expected: []string{"json"}}
	}

	return

}

func (p *parser) parseContent() (exp *ContentExpression, err error) {

	exp = &ContentExpression{}

	tok, lit, err := p.shouldBe(JSON, PARAM)
	if err != nil {
		return nil, &ParseError{Found: lit, Expected: []string{"json"}}
	}

	exp.Data, err = p.declare(tok, lit)
	if err != nil {
		return nil, &ParseError{Found: lit, Expected: []string{"json"}}
	}

	return

}