// Copyright 2018 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache-style
// license that can be found in the LICENSE file.

package logger_test

import (
	"os"
	
	"github.com/chai2010/logger"
)

func ExampleGetLogger() {
	var logger = logger.GetLogger()
	
	logger.SetLevel("DEBUG")
	logger.Debug("debug: ...")
	logger.Info("hello")
}

func ExampleNewStdLogger() {
	var logger = logger.NewStdLogger(os.Stderr)

	logger.Debug("debug: ...")
	logger.Info("hello")
}
