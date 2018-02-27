// Copyright 2018 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"log"
	"os"
	"runtime"

	"github.com/chai2010/logger"
)

func init() {
	logger.SetLogger(logger.NewStdLogger(os.Stderr, "", "", log.Lshortfile))
}

func main() {
	var logger = logger.GetLogger()

	logger.SetLevel("DEBUG")
	logger.Debug(runtime.Version())
	logger.Info("hello")
}
