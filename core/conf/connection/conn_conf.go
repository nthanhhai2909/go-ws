package connection

import (
	"net/http"
	"time"
)

type Configuration interface {
	GetHandshakeTimeout() time.Duration
	GetReadBufferSize() int
	GetWriteBufferSize() int
	GetSubProtocols() []string
	GetError() func(w http.ResponseWriter, r *http.Request, status int, reason error)
	GetCheckOrigin() func(r *http.Request) bool
	GetEnableCompression() bool
}
