package core

import (
	"github.com/gorilla/websocket"
	"log"
	"mem-ws/core/channel/inbound"
	"mem-ws/core/conf"
	"mem-ws/core/wserror"
)

type WebsocketConnectionConfigurationError struct {
	message string
}

func (e WebsocketConnectionConfigurationError) Error() string { return e.message }

type WebsocketConnectionFactory struct {
	upgrader       *websocket.Upgrader
	inboundChannel inbound.Channel[interface{}]
}

func NewWebSocketConnectionFactory(configuration conf.WebsocketConnectionConfiguration) (*WebsocketConnectionFactory, error) {
	upgrader, err := initWebsocketUpgrader(configuration)

	if err != nil {
		log.Panic("Invalid UpgradeConfiguration")
		return nil, err
	}

	return &WebsocketConnectionFactory{
		upgrader: upgrader,
		// TODO HGA WILL ADAPT TO CREATE BY CONFIGURATION
		inboundChannel: inbound.NewSubscribableChannel(),
	}, nil
}

func (factory *WebsocketConnectionFactory) GetUpgrader() *websocket.Upgrader {
	return factory.upgrader
}

func (factory *WebsocketConnectionFactory) GetInboundChannel() inbound.Channel[interface{}] {
	return factory.inboundChannel
}

func initWebsocketUpgrader(configuration conf.WebsocketConnectionConfiguration) (*websocket.Upgrader, error) {
	if configuration.GetReadBufferSize() <= 0 || configuration.GetWriteBufferSize() <= 0 {
		return nil, wserror.InvalidConfigurationError()
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
