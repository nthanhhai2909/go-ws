package stomp

import (
	"log"
	"mem-ws/socket"
	"mem-ws/socket/adapter/native"
	"mem-ws/socket/core/conf"
	"mem-ws/socket/core/stomp/registry"
	"net/http"
)

type Starter struct {
	factory  *WebsocketConnectionFactory
	registry *registry.EndpointRegistry
}

func NewWSStarter(configuration conf.WebsocketConnectionConfiguration) socket.WSStarter {
	factory, err := NewWebSocketConnectionFactory(configuration)
	if err != nil {
		log.Panic("Configuration Error")
	}

	return &Starter{
		factory:  factory,
		registry: &registry.EndpointRegistry{},
	}
}

func (starter *Starter) Handler(w http.ResponseWriter, r *http.Request) {
	factory := starter.factory
	conn, err := factory.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}
	websocketHandler := factory.SubProtocolWebsocketHandler
	// TODO ALLOW SETUP TEXT AND BINARY SIZE
	websocketSession := native.NewWebsocketSession(conn, 1024, 1024)
	err = websocketHandler.AfterConnectionEstablished(websocketSession)
	if err != nil {
		log.Println("Create connection error:", err)
		return
	}

	defer func() {
		// TODO HGA WILL PROCESS CLOSE STATUS LATER
		err = websocketHandler.AfterConnectionClosed(websocketSession, socket.Normal)
		if err != nil {
			log.Println("Error when close connection")
		}
	}()
	for {
		messageType, payload, err := conn.ReadMessage()
		// TODO HGA WILL HANDLE ERROR
		if err != nil {
			log.Println(err.Error())
			log.Println("starter: Error when send message")
			return
		}
		websocketMessage := socket.ToWebsocketMessage(messageType, payload)
		err = websocketHandler.HandleMessageFromClient(websocketSession, websocketMessage)
		if err != nil {
			log.Println("Error when send smsg")
			return
		}
	}
}

func (starter *Starter) AddEndPoint(endpoint string, handler socket.MessageHandler) {
}
