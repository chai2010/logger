- *赞助 BTC: 1Cbd6oGAUUyBi7X7MaR4np4nTmQZXVgkCW*
- *赞助 ETH: 0x623A3C3a72186A6336C79b18Ac1eD36e1c71A8a6*

----

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
