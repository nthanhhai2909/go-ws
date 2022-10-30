package socket

import (
	"github.com/gorilla/websocket"
	"log"
	"mem-ws/socket/channel"
	"mem-ws/socket/conf"
	"mem-ws/socket/stomp"
	"mem-ws/socket/wserror"
)

type WebsocketConnectionConfigurationError struct {
	message string
}

func (e WebsocketConnectionConfigurationError) Error() string { return e.message }

type WebsocketConnectionFactory struct {
	upgrader                    *websocket.Upgrader
	subProtocolWebsocketHandler WebsocketHandler
}

func NewWebSocketConnectionFactory(configuration conf.WebsocketConnectionConfiguration) (*WebsocketConnectionFactory, error) {
	upgrader, err := initWebsocketUpgrader(configuration)

	if err != nil {
		log.Panic("Invalid UpgradeConfiguration")
		return nil, err
	}
	// TODO HGA WILL ADAPT TO CREATE BY CONFIGURATION
	return &WebsocketConnectionFactory{
		subProtocolWebsocketHandler: &SubProtocolWebsocketHandler{
			ClientInboundChannel: channel.NewSubscribableChannel(),
			Sessions:             make(map[string]WebsocketSession),
			// TODO INIT BASED ON CONFIGURATION
			SubProtocolHandler: &stomp.SubProtocolHandler{},
		},
		upgrader: upgrader,
	}, nil
}

func (factory *WebsocketConnectionFactory) GetUpgrader() *websocket.Upgrader {
	return factory.upgrader
}

func (factory *WebsocketConnectionFactory) GetSubProtocolWebsocketHandler() WebsocketHandler {
	return factory.subProtocolWebsocketHandler
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
