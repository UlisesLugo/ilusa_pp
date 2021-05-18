// Code generated by gocc; DO NOT EDIT.

package token

import (
	"bytes"
	"fmt"
	"strconv"
	"unicode/utf8"
)

type Token struct {
	Type
	Lit []byte
	Pos
}

type Type int

const (
	INVALID Type = iota
	EOF
)

type Pos struct {
	Offset  int
	Line    int
	Column  int
	Context Context
}

func (p Pos) String() string {
	// If the context provides a filename, provide a human-readable File:Line:Column representation.
	switch src := p.Context.(type) {
	case Sourcer:
		return fmt.Sprintf("%s:%d:%d", src.Source(), p.Line, p.Column)
	default:
		return fmt.Sprintf("Pos(offset=%d, line=%d, column=%d)", p.Offset, p.Line, p.Column)
	}
}

type TokenMap struct {
	typeMap []string
	idMap   map[string]Type
}

func (m TokenMap) Id(tok Type) string {
	if int(tok) < len(m.typeMap) {
		return m.typeMap[tok]
	}
	return "unknown"
}

func (m TokenMap) Type(tok string) Type {
	if typ, exist := m.idMap[tok]; exist {
		return typ
	}
	return INVALID
}

func (m TokenMap) TokenString(tok *Token) string {
	return fmt.Sprintf("%s(%d,%s)", m.Id(tok.Type), tok.Type, tok.Lit)
}

func (m TokenMap) StringType(typ Type) string {
	return fmt.Sprintf("%s(%d)", m.Id(typ), typ)
}

// Equals returns returns true if the token Type and Lit are matches.
func (t *Token) Equals(rhs interface{}) bool {
	switch rhsT := rhs.(type) {
	case *Token:
		return t == rhsT || (t.Type == rhsT.Type && bytes.Equal(t.Lit, rhsT.Lit))
	default:
		return false
	}
}

// CharLiteralValue returns the string value of the char literal.
func (t *Token) CharLiteralValue() string {
	return string(t.Lit[1 : len(t.Lit)-1])
}

// Float32Value returns the float32 value of the token or an error if the token literal does not
// denote a valid float32.
func (t *Token) Float32Value() (float32, error) {
	if v, err := strconv.ParseFloat(string(t.Lit), 32); err != nil {
		return 0, err
	} else {
		return float32(v), nil
	}
}

// Float64Value returns the float64 value of the token or an error if the token literal does not
// denote a valid float64.
func (t *Token) Float64Value() (float64, error) {
	return strconv.ParseFloat(string(t.Lit), 64)
}

// IDValue returns the string representation of an identifier token.
func (t *Token) IDValue() string {
	return string(t.Lit)
}

// Int32Value returns the int32 value of the token or an error if the token literal does not
// denote a valid float64.
func (t *Token) Int32Value() (int32, error) {
	if v, err := strconv.ParseInt(string(t.Lit), 10, 64); err != nil {
		return 0, err
	} else {
		return int32(v), nil
	}
}

// Int64Value returns the int64 value of the token or an error if the token literal does not
// denote a valid float64.
func (t *Token) Int64Value() (int64, error) {
	return strconv.ParseInt(string(t.Lit), 10, 64)
}

// UTF8Rune decodes the UTF8 rune in the token literal. It returns utf8.RuneError if
// the token literal contains an invalid rune.
func (t *Token) UTF8Rune() (rune, error) {
	r, _ := utf8.DecodeRune(t.Lit)
	if r == utf8.RuneError {
		err := fmt.Errorf("Invalid rune")
		return r, err
	}
	return r, nil
}

// StringValue returns the string value of the token literal.
func (t *Token) StringValue() string {
	return string(t.Lit[1 : len(t.Lit)-1])
}

var TokMap = TokenMap{
	typeMap: []string{
		"INVALID",
		"$",
		"Program",
		"id",
		":",
		"empty",
		"main",
		"(",
		")",
		"{",
		"}",
		"Class",
		"atributes",
		"inherits",
		"methods",
		"public",
		"private",
		"protected",
		";",
		"variables",
		"[",
		"cte_i",
		"]",
		",",
		"int",
		"char",
		"float",
		"function",
		"body",
		"void",
		"=",
		"return",
		"input",
		"output",
		"cte_string",
		"if",
		"else",
		"while",
		"for",
		"+",
		"-",
		"*",
		"/",
		"cte_float",
		"&&",
		"||",
		"!=",
		"<",
		">",
		">=",
		"<=",
		"==",
	},

	idMap: map[string]Type{
		"INVALID":    0,
		"$":          1,
		"Program":    2,
		"id":         3,
		":":          4,
		"empty":      5,
		"main":       6,
		"(":          7,
		")":          8,
		"{":          9,
		"}":          10,
		"Class":      11,
		"atributes":  12,
		"inherits":   13,
		"methods":    14,
		"public":     15,
		"private":    16,
		"protected":  17,
		";":          18,
		"variables":  19,
		"[":          20,
		"cte_i":      21,
		"]":          22,
		",":          23,
		"int":        24,
		"char":       25,
		"float":      26,
		"function":   27,
		"body":       28,
		"void":       29,
		"=":          30,
		"return":     31,
		"input":      32,
		"output":     33,
		"cte_string": 34,
		"if":         35,
		"else":       36,
		"while":      37,
		"for":        38,
		"+":          39,
		"-":          40,
		"*":          41,
		"/":          42,
		"cte_float":  43,
		"&&":         44,
		"||":         45,
		"!=":         46,
		"<":          47,
		">":          48,
		">=":         49,
		"<=":         50,
		"==":         51,
	},
}
