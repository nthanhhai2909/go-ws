package stomp

import (
	"github.com/gorilla/websocket"
	"log"
	"mem-ws/core/conf"
	"mem-ws/core/errors"
	"mem-ws/native/handler"
	"mem-ws/native/session"
)

type WebsocketConnectionConfigurationError struct {
	message string
}

func (e WebsocketConnectionConfigurationError) Error() string { return e.message }

type WebsocketConnectionFactory struct {
	Upgrader                    *websocket.Upgrader
	SubProtocolWebsocketHandler handler.IWebsocketHandler
}

func NewWebSocketConnectionFactory(configuration conf.WebsocketConnectionConfiguration) (*WebsocketConnectionFactory, error) {
	upgrader, err := initWebsocketUpgrader(configuration)

	if err != nil {
		log.Panic("Invalid UpgradeConfiguration")
		return nil, err
	}
	// TODO HGA WILL ADAPT TO CREATE BY CONFIGURATION
	return &WebsocketConnectionFactory{
		SubProtocolWebsocketHandler: &handler.NativeSubProtocolHandler{
			Sessions: make(map[string]session.ISession),
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
