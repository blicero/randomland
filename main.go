// /home/krylon/go/src/github.com/blicero/randomland/main.go
// -*- mode: go; coding: utf-8; -*-
// Created on 26. 03. 2025 by Benjamin Walkenhorst
// (c) 2025 Benjamin Walkenhorst
// Time-stamp: <2025-03-26 18:19:03 krylon>

package main

import (
	"fmt"
	"os"

	"github.com/blicero/randomland/common"
	"github.com/blicero/randomland/shell"

	_ "github.com/blicero/krylib" // To keep the dependency for the build script
)

func main() {
	fmt.Printf("%s %s, built on %s\n",
		common.AppName,
		common.Version,
		common.BuildStamp.Format(common.TimestampFormatMinute))

	var (
		err error
		s   *shell.Shell
	)

	if s, err = shell.New(); err != nil {
		fmt.Fprintf(
			os.Stderr,
			"Error creating Shell: %s\n",
			err.Error())
		os.Exit(1)
	} else if err = s.Run(); err != nil {
		fmt.Fprintf(
			os.Stderr,
			"Error running Shell: %s\n",
			err.Error())
		os.Exit(1)
	}
}
