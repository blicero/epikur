// /home/krylon/go/src/github.com/blicero/epikur/main.go
// -*- mode: go; coding: utf-8; -*-
// Created on 03. 03. 2025 by Benjamin Walkenhorst
// (c) 2025 Benjamin Walkenhorst
// Time-stamp: <2025-03-03 19:10:58 krylon>

package main

import (
	"fmt"

	"github.com/blicero/epikur/common"
)

func main() {
	fmt.Printf("%s %s\n(c) 2025 Benjamin Walkenhorst\n\n",
		common.AppName,
		common.Version)

	fmt.Println("Nothing to see here, move along.")
}
