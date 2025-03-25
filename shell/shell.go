// /home/krylon/go/src/github.com/blicero/randomland/shell/shell.go
// -*- mode: go; coding: utf-8; -*-
// Created on 25. 03. 2025 by Benjamin Walkenhorst
// (c) 2025 Benjamin Walkenhorst
// Time-stamp: <2025-03-25 17:02:29 krylon>

// Package shell provides the text-based user interface.
package shell

import (
	"log"

	"github.com/blicero/krylib"
	"github.com/blicero/randomland/common"
	"github.com/blicero/randomland/logdomain"
	"github.com/c-bata/go-prompt"
)

var commands = []prompt.Suggest{
	{Text: "look", Description: "Inspect your surroundings"},
	{Text: "take", Description: "Take an object"},
	{Text: "attack", Description: "Attack a monster"},
	{Text: "go", Description: "Go to a different location"},
	{Text: "save", Description: "Save the game"},
	{Text: "load", Description: "Load a saved game"},
	{Text: "quit", Description: "Quit"},
}

func makeCompleter(sugg []prompt.Suggest) func(d prompt.Document) []prompt.Suggest {
	return func(d prompt.Document) []prompt.Suggest {
		return prompt.FilterHasPrefix(sugg, d.GetWordBeforeCursor(), true)
	}
} // func makeCompleter(sugg []prompt.Suggest) func(d prompt.Document) []prompt.Suggest

// Shell provides a text-based interface to the user.
type Shell struct {
	log *log.Logger
}

// New creates a new Shell
func New() (*Shell, error) {
	var (
		err error
		s   = new(Shell)
	)

	if s.log, err = common.GetLogger(logdomain.Shell); err != nil {
		return nil, err
	}

	return s, nil
} // func New() (*Shell, error)

func (s *Shell) Run() error {
	return krylib.ErrNotImplemented
}
