package stomp

import (
	"log"
	"mem-ws/core"
	"mem-ws/core/conf"
	"mem-ws/core/subprotocols/stomp/registry"
	error2 "mem-ws/native/enums"
	"mem-ws/native/message"
	"mem-ws/native/session"
	"net/http"
)

type Starter struct {
	factory  *WebsocketConnectionFactory
	registry *registry.EndpointRegistry
}

func NewWSStarter(configuration conf.WebsocketConnectionConfiguration) core.WSStarter {
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
	websocketSession := session.NewWebsocketSession(conn, 1024, 1024)
	err = websocketHandler.AfterConnectionEstablished(websocketSession)
	if err != nil {
		log.Println("Create connection enums:", err)
		return
	}

	defer func() {
		// TODO HGA WILL PROCESS CLOSE STATUS LATER
		err = websocketHandler.AfterConnectionClosed(websocketSession, error2.Normal)
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
		websocketMessage := message.ToMessage(messageType, payload)
		err = websocketHandler.HandleMessageFromClient(websocketSession, websocketMessage)
		if err != nil {
			log.Println("Error when send smsg")
			return
		}
	}
}

func (starter *Starter) AddEndPoint(endpoint string, handler core.MessageHandler) {
}
