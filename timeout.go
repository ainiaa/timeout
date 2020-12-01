package timeout

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var to *Timeout

// Option for timeout
type Option func(*Timeout)

// WithTimeout set timeout
func WithTimeout(timeout time.Duration) Option {
	return func(t *Timeout) {
		t.timeout = timeout
	}
}

// WithHandler add gin handler
func WithHandler(h gin.HandlerFunc) Option {
	return func(t *Timeout) {
		t.handler = h
	}
}

// WithResponse add gin handler
func WithResponse(h gin.HandlerFunc) Option {
	return func(t *Timeout) {
		t.response = h
	}
}

func WithVersion(version string) Option {
	return func(t *Timeout) {
		t.version = version
	}
}

func defaultResponse(c *gin.Context) {
	c.JSON(http.StatusRequestTimeout, Response{
		Code:       http.StatusRequestTimeout,
		Success:    false,
		ServerTime: time.Now().Unix(),
		Message:    http.ErrHandlerTimeout.Error(),
		Version:    to.version,
		Timestamp:  time.Now().Format("2006-01-02 15:04:05"),
		ExecTime:   0,
	})
}

// Timeout struct
type Timeout struct {
	timeout  time.Duration
	handler  gin.HandlerFunc
	response gin.HandlerFunc
	version  string
}

// New wraps a handler and aborts the process of the handler if the timeout is reached
func New(opts ...Option) gin.HandlerFunc {
	const (
		defaultTimeout = 5 * time.Second
	)

	to = &Timeout{
		timeout:  defaultTimeout,
		handler:  nil,
		response: defaultResponse,
	}

	// Loop through each option
	for _, opt := range opts {
		// Call the option giving the instantiated
		opt(to)
	}

	if to.timeout <= 0 {
		return to.handler
	}
	if to.handler == nil {
		to.handler = func(c *gin.Context) {
			time.Sleep(to.timeout + 1*time.Second)
		}
	}

	return func(c *gin.Context) {
		ch := make(chan struct{}, 1)

		go func() {
			defer func() {
				_ = gin.Recovery()
			}()
			to.handler(c)
			ch <- struct{}{}
		}()

		select {
		case <-ch:
			c.Next()
		case <-time.After(to.timeout):
			c.Abort()
			to.response(c)
			return
		}
	}
}
