package provider

import (
	"log"
	"mem-ws/socket"
	"mem-ws/socket/adapter/native"
	"net/http"
)

type WSStarter interface {
	Handler(w http.ResponseWriter, r *http.Request)
}

type wsStarter struct {
	factory *WebsocketConnectionFactory
}

func NewWSStarter(configuration WebsocketConnectionConfiguration) WSStarter {
	factory, err := NewWebSocketConnectionFactory(configuration)
	if err != nil {
		log.Panic("Configuration Error")
	}

	return &wsStarter{
		factory: factory,
	}
}

func (starter *wsStarter) Handler(w http.ResponseWriter, r *http.Request) {
	conn, err := starter.factory.GetUpgrader().Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}
	websocketHandler := starter.factory.GetSubProtocolWebsocketHandler()
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
