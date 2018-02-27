# Simple level logging package.

```go
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
```

Output:

```
hello.go:25: [DEBUG] go1.10
hello.go:26: [INFO] hello
```
