package provider

import (
	"net/http"
	"time"
)

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
