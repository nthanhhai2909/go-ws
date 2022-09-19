package conf

import (
	"net/http"
	"time"
)

const (
	ReadBufferSize  = 1024
	WriteBufferSize = 1024
	Endpoint        = "/ws"
)

type WebsocketConnectionConfiguration interface {
	GetHandshakeTimeout() time.Duration
	GetReadBufferSize() int
	GetWriteBufferSize() int
	GetSubProtocols() []string
	GetError() func(w http.ResponseWriter, r *http.Request, status int, reason error)
	GetCheckOrigin() func(r *http.Request) bool
	GetEnableCompression() bool
}

func NewDefaultWebsocketConnectionConfiguration() *DefaultWebsocketConnectionConfiguration {
	return &DefaultWebsocketConnectionConfiguration{}
}

type DefaultWebsocketConnectionConfiguration struct {
}

func (conf *DefaultWebsocketConnectionConfiguration) GetHandshakeTimeout() time.Duration {
	return time.Duration(0)
}

func (conf *DefaultWebsocketConnectionConfiguration) GetReadBufferSize() int {
	return ReadBufferSize
}

func (conf *DefaultWebsocketConnectionConfiguration) GetWriteBufferSize() int {
	return WriteBufferSize
}

func (conf *DefaultWebsocketConnectionConfiguration) GetSubProtocols() []string {
	return make([]string, 0)
}

func (conf *DefaultWebsocketConnectionConfiguration) GetError() func(w http.ResponseWriter, r *http.Request, status int, reason error) {
	return nil
}

func (conf *DefaultWebsocketConnectionConfiguration) GetCheckOrigin() func(r *http.Request) bool {
	return nil
}

func (conf *DefaultWebsocketConnectionConfiguration) GetEnableCompression() bool {
	return false
}
