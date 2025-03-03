// /home/krylon/go/src/github.com/blicero/epikur/grammar/01_parser_test.go
// -*- mode: go; coding: utf-8; -*-
// Created on 03. 03. 2025 by Benjamin Walkenhorst
// (c) 2025 Benjamin Walkenhorst
// Time-stamp: <2025-03-03 20:44:26 krylon>

package grammar

import (
	"testing"

	"github.com/alecthomas/participle/v2"
	"github.com/davecgh/go-spew/spew"
)

var par *participle.Parser[Value]

func TestBuildParser(t *testing.T) {
	par = New()
} // func TestBuildParser(t *testing.T)

func TestPrimitive(t *testing.T) {
	if par == nil {
		t.SkipNow()
	}

	type testCase struct {
		name   string
		input  string
		output Value
		err    bool
	}

	var samples = []testCase{
		{
			name:   "int",
			input:  "5",
			output: &Integer{Val: 5},
		},
		{
			name:   "string",
			input:  `"Hello there"`,
			output: &String{Val: "Hello there"},
		},
		{
			name:   "real",
			input:  "3.141592653589793",
			output: &Real{Val: 3.141592653589793},
		},
		{
			name:   "array",
			input:  "[1 2 3]",
			output: &Array{Values: []Value{&Integer{Val: 1}, &Integer{Val: 2}, &Integer{Val: 3}}},
		},
	}

	for _, c := range samples {
		var (
			err error
			v   *Value
		)

		if v, err = par.ParseString(c.name, c.input); err != nil {
			if !c.err {
				t.Errorf("Failed to parse %s: %s",
					c.name,
					err.Error())
			}
		} else if !(*v).Equal(c.output) {
			t.Errorf("Unexpected result from parsing %s : %q (expected %q)",
				c.name,
				spew.Sdump(v),
				spew.Sdump(c.output))
		}
	}
} // func TestPrimitive(t *testing.T)
