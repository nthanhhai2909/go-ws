package ws

import (
	"github.com/gorilla/websocket"
	"log"
	"mem-ws/wscore"
	"net/http"
)

type Service interface {
	Handler(w http.ResponseWriter, r *http.Request)
}

type service struct {
	upg websocket.Upgrader
	hub wscore.Hub
}

func New(upg websocket.Upgrader, hub wscore.Hub) Service {
	return &service{
		upg: upg,
		hub: hub,
	}
}

func (s *service) Handler(w http.ResponseWriter, r *http.Request) {
	conn, err := s.upg.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	client := &wscore.Client{Conn: conn, Outbound: make(chan []byte)}
	s.hub.Register <- client
	defer func() {
		conn.Close()
		s.hub.Unregister <- client
	}()
	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = conn.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}

}
