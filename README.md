# timeout
timeout gin framework middleware

thank to github.com/gin-contrib/timeout
## Example

```go
package main

import (
	"net/http"
	"time"

	"github.com/ainiaa/timeout"
	"github.com/gin-gonic/gin"
)

func emptyHandler(c *gin.Context) {
	time.Sleep(200 * time.Microsecond)
	c.String(http.StatusOK, "")
}

func main() {
	r := gin.New()

	r.GET("/", timeout.New(
		timeout.WithTimeout(100*time.Microsecond),
	))

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
```

### custom error response

Add new error response func:

```go
func testResponse(c *gin.Context) {
	c.String(http.StatusRequestTimeout, "test response")
}
```

Add `WithResponse` option.

```go
	r.GET("/", timeout.New(
		WithTimeout(100*time.Microsecond),
		WithHandler(emptyHandler),
		WithResponse(testResponse),
	))
```