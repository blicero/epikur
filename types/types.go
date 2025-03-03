// /home/krylon/go/src/github.com/blicero/epikur/types/types.go
// -*- mode: go; coding: utf-8; -*-
// Created on 03. 03. 2025 by Benjamin Walkenhorst
// (c) 2025 Benjamin Walkenhorst
// Time-stamp: <2025-03-03 19:08:24 krylon>

// Packages types provides symbolic constants to refer to the primitive
// types of the Epikur programming language.
package types

//go:generate stringer -type=ID

// ID identifies a type.
type ID uint8

const (
	Null ID = iota
	Integer
	Real
	String
	Array
	Map
	Object
	IO
	Proc
)
