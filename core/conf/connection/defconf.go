package connection

import (
	"net/http"
	"time"
)

func NewDefaultConfiguration() *DefaultConfiguration {
	return &DefaultConfiguration{}
}

type DefaultConfiguration struct {
}

func (conf *DefaultConfiguration) GetHandshakeTimeout() time.Duration {
	return time.Duration(0)
}

func (conf *DefaultConfiguration) GetReadBufferSize() int {
	return ReadBufferSize
}

func (conf *DefaultConfiguration) GetWriteBufferSize() int {
	return WriteBufferSize
}

func (conf *DefaultConfiguration) GetSubProtocols() []string {
	return make([]string, 0)
}

func (conf *DefaultConfiguration) GetError() func(w http.ResponseWriter, r *http.Request, status int, reason error) {
	return nil
}

func (conf *DefaultConfiguration) GetCheckOrigin() func(r *http.Request) bool {
	return nil
}

func (conf *DefaultConfiguration) GetEnableCompression() bool {
	return false
}
