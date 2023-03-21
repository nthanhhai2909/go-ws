package core

import (
	"github.com/gorilla/websocket"
	"log"
	"mem-ws/core/conf"
	"mem-ws/core/conf/connection"
	"mem-ws/core/errors"
	"mem-ws/core/subprotocols/stomp"
	"mem-ws/native/handler"
	"sync"
)

type WebsocketConnectionConfigurationError struct {
	message string
}

// TODO SUPPORT CONFIGURATION SUB-PROTOCOL - CURRENT HARDCODE STOMP
func (e WebsocketConnectionConfigurationError) Error() string { return e.message }

type WebsocketConnectionFactory struct {
	Upgrader         *websocket.Upgrader
	WebsocketHandler handler.IWebsocketHandler
}

func NewWebSocketConnectionFactory(conf conf.Configuration) (*WebsocketConnectionFactory, error) {
	upgrader, err := initWebsocketUpgrader(conf.ConnectionConf)

	if err != nil {
		log.Panic("Invalid UpgradeConfiguration")
		return nil, err
	}
	// TODO HGA WILL ADAPT TO CREATE BY CONFIGURATION
	// CURRENTLY ONLY SUPPORT STOMP
	// TODO USE conf.BROKER
	WebsocketHandler := &handler.NativeWebsocketHandler{
		Sessions:           sync.Map{},
		SubProtocolHandler: stomp.NewProtocolHandler(conf.BrokerRegistry.StompBrokerRegistration),
	}
	return &WebsocketConnectionFactory{
		WebsocketHandler: WebsocketHandler,
		Upgrader:         upgrader,
	}, nil
}

func initWebsocketUpgrader(configuration connection.Configuration) (*websocket.Upgrader, error) {
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
