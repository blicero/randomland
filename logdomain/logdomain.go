// /home/krylon/go/src/github.com/blicero/krylisp/logdomain/logdomain.go
// -*- mode: go; coding: utf-8; -*-
// Created on 21. 12. 2024 by Benjamin Walkenhorst
// (c) 2024 Benjamin Walkenhorst
// Time-stamp: <2025-03-24 12:08:44 krylon>

package logdomain

//go:generate stringer -type=ID

type ID uint8

const (
	Engine ID = iota
)

func AllDomains() []ID {
	return []ID{
		Engine,
	}
} // func AllDomains() []ID
