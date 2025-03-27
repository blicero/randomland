// /home/krylon/go/src/github.com/blicero/randomland/shell/shell.go
// -*- mode: go; coding: utf-8; -*-
// Created on 25. 03. 2025 by Benjamin Walkenhorst
// (c) 2025 Benjamin Walkenhorst
// Time-stamp: <2025-03-27 15:47:49 krylon>

// Package shell provides the text-based user interface.
package shell

import (
	"fmt"
	"log"
	"os"

	"github.com/anmitsu/go-shlex"
	"github.com/blicero/krylib"
	"github.com/blicero/randomland/common"
	"github.com/blicero/randomland/logdomain"
	"github.com/c-bata/go-prompt"
)

const prefix = ">>> "

var commands = []prompt.Suggest{
	{Text: "look", Description: "Inspect your surroundings"},
	{Text: "take", Description: "Take an object"},
	{Text: "attack", Description: "Attack a monster"},
	{Text: "go", Description: "Go to a different location"},
	{Text: "save", Description: "Save the game"},
	{Text: "load", Description: "Load a saved game"},
	{Text: "quit", Description: "Quit"},
}

// type completer func(d prompt.Document) []prompt.Suggest

func makeCompleter(sugg []prompt.Suggest) func(d prompt.Document) []prompt.Suggest {
	return func(d prompt.Document) []prompt.Suggest {
		return prompt.FilterHasPrefix(sugg, d.GetWordBeforeCursor(), true)
	}
} // func makeCompleter(sugg []prompt.Suggest) completer

type shellCmd func(cmd []string) error

// Shell provides a text-based interface to the user.
type Shell struct {
	log    *log.Logger
	cmpl   func(d prompt.Document) []prompt.Suggest
	prompt *prompt.Prompt
	cmdmap map[string]shellCmd
	round  int
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

	s.cmpl = makeCompleter(commands)
	s.cmdmap = map[string]shellCmd{
		"attack": s.doAttack,
		"go":     s.doGo,
		"look":   s.doLook,
		"take":   s.doTake,
		"load":   s.doLoad,
		"save":   s.doSave,
		"quit":   s.doQuit,
	}

	s.prompt = prompt.New(
		s.executor,
		s.cmpl,
		prompt.OptionPrefix(prefix),
	)

	s.round++

	return s, nil
} // func New() (*Shell, error)

// Run executes the Shell's main loop
func (s *Shell) Run() error {
	var (
		err error
	)

	defer func() {
		if x := recover(); x != nil {
			s.log.Printf("[ERROR] Panic while running prompt: %s\n",
				x)
		}
	}()

	s.prompt.Run()

	return err
} // func (s *Shell) Run() error

func (s *Shell) executor(in string) {
	s.log.Printf("[TRACE] Attempting to execute %s\n",
		in)

	var (
		err     error
		handler shellCmd
		pieces  []string
		ok      bool
	)

	if pieces, err = shlex.Split(in, false); err != nil {
		s.log.Printf("[ERROR] Cannot tokenize input %q: %s\n",
			in,
			err.Error())
	}

	if handler, ok = s.cmdmap[pieces[0]]; !ok {
		fmt.Printf("Unknown command: %s\n", pieces[0])
		return
	} else if err = handler(pieces); err != nil {
		s.log.Printf("[ERROR] Error handling %s: %s\n",
			pieces[0],
			err.Error())
	}
} // func (s *Shell) executor(in string)

func (s *Shell) doAttack(cmd []string) error {
	return krylib.ErrNotImplemented
} // func (s *Shell) doAttack(cmd []string) error

func (s *Shell) doGo(cmd []string) error {
	return krylib.ErrNotImplemented
} // func (s *Shell) doGo(cmd []string) error

func (s *Shell) doLook(cmd []string) error {
	return krylib.ErrNotImplemented
} // func (s *Shell) doLook(cmd []string) error

func (s *Shell) doTake(cmd []string) error {
	return krylib.ErrNotImplemented
} // func (s *Shell) doTake(cmd []string) error

func (s *Shell) doSave(cmd []string) error {
	return krylib.ErrNotImplemented
} // func (s *Shell) doSave(cmd []string) error

func (s *Shell) doLoad(cmd []string) error {
	return krylib.ErrNotImplemented
} // func (s *Shell) doLoad(cmd []string) error

func (s *Shell) doQuit(cmd []string) error {
	fmt.Println("So long, and thanks for all the random numbers!")
	os.Exit(0)
	return nil // The linter barked at the missing return value. *shrug*
} // func (s *Shell) doQuit(cmd []string) error
