// /home/krylon/go/src/github.com/blicero/epikur/grammar/grammar.go
// -*- mode: go; coding: utf-8; -*-
// Created on 03. 03. 2025 by Benjamin Walkenhorst
// (c) 2025 Benjamin Walkenhorst
// Time-stamp: <2025-03-03 20:35:34 krylon>

package grammar

import (
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/blicero/epikur/types"
)

var lex = lexer.MustSimple([]lexer.SimpleRule{
	{Name: "Real", Pattern: `(\d+[.]\d+(?:e(\d+))?)`},
	{Name: "Integer", Pattern: `\d+|0x([0-9a-f]+)`},
	{Name: "OpenParen", Pattern: `\(`},
	{Name: `CloseParen`, Pattern: `\)`},
	{Name: `OpenBrace`, Pattern: `\{`},
	{Name: `CloseBrace`, Pattern: `\}`},
	{Name: `OpenBracket`, Pattern: `\[`},
	{Name: `CloseBracket`, Pattern: `\]`},
	{Name: `Comma`, Pattern: `,`},
	{Name: "String", Pattern: `"(?:[^\"]*)"`},
	{Name: "Operator", Pattern: `[-+*/%<>!&|]`},
	{Name: "Semicolon", Pattern: ";"},
	{Name: "Colon", Pattern: ":"},
	{Name: "Hash", Pattern: "#"},
	{Name: "Dot", Pattern: "[.]"},
	{Name: "Underscore", Pattern: "_"},
	{Name: "Whitespace", Pattern: `\s+`},
})

// New creates a Parser
func New() *participle.Parser[Value] {
	par := participle.MustBuild[Value](
		participle.Lexer(lex),
		participle.Unquote("String"),
		participle.Elide("Whitespace"),
		participle.Union[Value](&Integer{}, &Real{}, &String{}, &Array{}),
	)

	return par
}

// Value is a common interface for all the primitive data types Epikur includes.
type Value interface {
	Type() types.ID
	Equal(other Value) bool
}

// Integer is a whole number.
type Integer struct {
	Val int64 `parser:"@Integer"`
}

// Type returns the Type ID of the receiver
func (i *Integer) Type() types.ID { return types.Integer }

// Equal returns if true if the receiver and the argument have equal values.
func (i *Integer) Equal(other Value) bool {
	switch v := other.(type) {
	case *Integer:
		return i.Val == v.Val
	case *Real:
		return float64(i.Val) == v.Val
	default:
		return false
	}
} // func (i *Integer) Equal(other Value) bool

// Real is real number
type Real struct {
	Val float64 `parser:"@Real"`
}

// Type returns the Type ID of the receiver
func (r *Real) Type() types.ID { return types.Real }

// Equal returns if true if the receiver and the argument have equal values.
func (r *Real) Equal(other Value) bool {
	switch v := other.(type) {
	case *Real:
		return r.Val == v.Val
	case *Integer:
		return r.Val == float64(v.Val)
	default:
		return false
	}
} // func (r *Real) Equal(other Value) bool

// String is a string, i.e. a chunk of text.
type String struct {
	Val string `parser:"@String"`
}

// Type returns the Type ID of the receiver
func (s *String) Type() types.ID { return types.String }

// Equal returns if true if the receiver and the argument have equal values.
func (s *String) Equal(other Value) bool {
	switch v := other.(type) {
	case *String:
		return s.Val == v.Val
	default:
		return false
	}
}

// Array is a sequence of Values
type Array struct {
	Values []Value `parser:"OpenBracket @@* CloseBracket"`
}

// Type returns the Type ID of the receiver
func (a *Array) Type() types.ID { return types.Array }

// Equal returns if true if the receiver and the argument have equal values.
func (a *Array) Equal(other Value) bool {
	switch v := other.(type) {
	case *Array:
		if len(a.Values) != len(v.Values) {
			return false
		}

		for idx, val := range a.Values {
			if !val.Equal(v.Values[idx]) {
				return false
			}
		}

		return true
	default:
		return false
	}
} // func (a *Array) Equal(other Value) bool
