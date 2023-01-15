package provider

import (
	"github.com/gorilla/websocket"
	"log"
	"mem-ws/socket"
	"mem-ws/socket/adapter/native"
	"mem-ws/socket/core/errors"
	"mem-ws/socket/core/stomp"
)

type WebsocketConnectionConfigurationError struct {
	message string
}

func (e WebsocketConnectionConfigurationError) Error() string { return e.message }

type WebsocketConnectionFactory struct {
	upgrader                    *websocket.Upgrader
	subProtocolWebsocketHandler socket.IWebsocketHandler
}

func NewWebSocketConnectionFactory(configuration WebsocketConnectionConfiguration) (*WebsocketConnectionFactory, error) {
	upgrader, err := initWebsocketUpgrader(configuration)

	if err != nil {
		log.Panic("Invalid UpgradeConfiguration")
		return nil, err
	}
	// TODO HGA WILL ADAPT TO CREATE BY CONFIGURATION
	return &WebsocketConnectionFactory{
		subProtocolWebsocketHandler: &native.SubProtocolWebsocketHandler{
			Sessions: make(map[string]socket.IWebsocketSession),
			// TODO SUPPORT INIT SUB-PROTOCOL BY CONFIGURATION
			SubProtocolHandler: stomp.NewProtocolHandler(),
		},
		upgrader: upgrader,
	}, nil
}

func (factory *WebsocketConnectionFactory) GetUpgrader() *websocket.Upgrader {
	return factory.upgrader
}

func (factory *WebsocketConnectionFactory) GetSubProtocolWebsocketHandler() socket.IWebsocketHandler {
	return factory.subProtocolWebsocketHandler
}

func initWebsocketUpgrader(configuration WebsocketConnectionConfiguration) (*websocket.Upgrader, error) {
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
