package wscore

import (
	"github.com/gorilla/websocket"
	"log"
	"mem-ws/wscore/conf"
	"mem-ws/wscore/errors"
)

type WebsocketConnectionConfigurationError struct {
	message string
}

func (e WebsocketConnectionConfigurationError) Error() string { return e.message }

type WebsocketConnectionFactory struct {
	upgrader *websocket.Upgrader
}

func NewWebSocketConnectionFactory(configuration conf.WebsocketConnectionConfiguration) (*WebsocketConnectionFactory, error) {
	upgrader, err := initWebsocketUpgrader(configuration)

	if err != nil {
		log.Panic("Invalid UpgradeConfiguration")
		return nil, err
	}

	return &WebsocketConnectionFactory{
		upgrader: upgrader,
	}, nil
}

func (factory *WebsocketConnectionFactory) GetUpgrader() *websocket.Upgrader {
	return factory.upgrader
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
