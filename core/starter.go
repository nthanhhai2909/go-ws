package core

import (
	"flag"
	"fmt"
	"log"
	"mem-ws/core/conf"
	"mem-ws/native/enums"
	"mem-ws/native/message"
	"mem-ws/native/session"
	"net/http"
)

var addr = flag.String("addr", "localhost:8999", "http service address")

type Starter struct {
	factory *WebsocketConnectionFactory
	conf    conf.Configuration
}

func NewWSStarter(conf conf.Configuration) *Starter {
	factory, err := NewWebSocketConnectionFactory(conf)
	if err != nil {
		log.Panic("Configuration Error")
	}

	return &Starter{
		factory: factory,
		conf:    conf,
	}
}

func (starter *Starter) Start() {
	http.HandleFunc(starter.conf.BrokerRegistry.ApplicationDestinationPrefix, starter.Handler)
	fmt.Println("Server start listening at: localhost:8999")
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func (starter *Starter) Handler(w http.ResponseWriter, r *http.Request) {
	factory := starter.factory
	conn, err := factory.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}
	websocketHandler := factory.WebsocketHandler
	// TODO ALLOW SETUP TEXT AND BINARY SIZE
	websocketSession := session.NewWebsocketSession(conn, 1024, 1024)
	err = websocketHandler.AfterConnectionEstablished(websocketSession)
	if err != nil {
		log.Println("Create connection enums:", err)
		return
	}

	defer func() {
		// TODO HGA WILL PROCESS CLOSE STATUS LATER
		err = websocketHandler.AfterConnectionClosed(websocketSession, enums.Normal)
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
