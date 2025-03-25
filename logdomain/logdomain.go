// /home/krylon/go/src/github.com/blicero/krylisp/logdomain/logdomain.go
// -*- mode: go; coding: utf-8; -*-
// Created on 21. 12. 2024 by Benjamin Walkenhorst
// (c) 2024 Benjamin Walkenhorst
// Time-stamp: <2025-03-25 16:58:50 krylon>

package logdomain

//go:generate stringer -type=ID

type ID uint8

const (
	Engine ID = iota
	Shell
)

func AllDomains() []ID {
	return []ID{
		Engine,
		Shell,
	}
} // func AllDomains() []ID
