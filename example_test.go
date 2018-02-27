// Copyright 2018 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache-style
// license that can be found in the LICENSE file.

package logger_test

import (
	"log"
	"os"

	"github.com/chai2010/logger"
)

func Example() {
	var logger = logger.GetLogger()

	logger.SetLevel("DEBUG")
	logger.Debug("1+1=2")
	logger.Info("hello")
}

func ExampleNewStdLogger() {
	var logger = logger.NewStdLogger(os.Stderr, "", log.Lshortfile)

	logger.Debug("1+1=2")
	logger.Info("hello")
}
