package provider

import (
	"log"
	"mem-ws/socket"
	"mem-ws/socket/adapter/native"
	"net/http"
)

type WSContainer interface {
	Handler(w http.ResponseWriter, r *http.Request)
}

type wscontainer struct {
	factory *WebsocketConnectionFactory
}

func NewWSContainer(websocketConnectionFactory *WebsocketConnectionFactory) WSContainer {
	return &wscontainer{
		factory: websocketConnectionFactory,
	}
}

func (container *wscontainer) Handler(w http.ResponseWriter, r *http.Request) {
	conn, err := container.factory.GetUpgrader().Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}
	websocketHandler := container.factory.GetSubProtocolWebsocketHandler()
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
			log.Println("container: Error when send message")
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
