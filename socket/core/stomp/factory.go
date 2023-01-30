package stomp

import (
	"github.com/gorilla/websocket"
	"log"
	"mem-ws/socket"
	"mem-ws/socket/adapter/native"
	"mem-ws/socket/core/conf"
	"mem-ws/socket/core/errors"
)

type WebsocketConnectionConfigurationError struct {
	message string
}

func (e WebsocketConnectionConfigurationError) Error() string { return e.message }

type WebsocketConnectionFactory struct {
	Upgrader                    *websocket.Upgrader
	SubProtocolWebsocketHandler socket.IWebsocketHandler
}

func NewWebSocketConnectionFactory(configuration conf.WebsocketConnectionConfiguration) (*WebsocketConnectionFactory, error) {
	upgrader, err := initWebsocketUpgrader(configuration)

	if err != nil {
		log.Panic("Invalid UpgradeConfiguration")
		return nil, err
	}
	// TODO HGA WILL ADAPT TO CREATE BY CONFIGURATION
	return &WebsocketConnectionFactory{
		SubProtocolWebsocketHandler: &native.SubProtocolWebsocketHandler{
			Sessions: make(map[string]socket.IWebsocketSession),
			// TODO SUPPORT INIT SUB-PROTOCOL BY CONFIGURATION
			SubProtocolHandler: NewProtocolHandler(),
		},
		Upgrader: upgrader,
	}, nil
}

func initWebsocketUpgrader(configuration conf.WebsocketConnectionConfiguration) (*websocket.Upgrader, error) {
	if configuration.GetReadBufferSize() <= 0 || configuration.GetWriteBufferSize() <= 0 {
		return nil, errors.InvalidConfigurationError()
	}

	upgrader := websocket.Upgrader{
		ReadBufferSize:    configuration.GetReadBufferSize(),
		WriteBufferSize:   configuration.GetWriteBufferSize(),
		Subprotocols:      configuration.GetSubProtocols(),
		Error:             configuration.GetError(),
		CheckOrigin:       configuration.GetCheckOrigin(),
		EnableCompression: configuration.GetEnableCompression(),
	}

	return &upgrader, nil
}
