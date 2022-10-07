package core

import (
	"log"
	"net/http"
)

type WSContainer interface {
	Handler(w http.ResponseWriter, r *http.Request)
}

type wscontainer struct {
	websocketConnectionFactory *WebsocketConnectionFactory
}

func NewWSContainer(websocketConnectionFactory *WebsocketConnectionFactory) WSContainer {
	return &wscontainer{
		websocketConnectionFactory: websocketConnectionFactory,
	}
}

func (container *wscontainer) Handler(w http.ResponseWriter, r *http.Request) {
	conn, err := container.websocketConnectionFactory.GetUpgrader().Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	handler, err := container.websocketConnectionFactory.GetInboundChannel().Connect(conn)
	if err != nil {
		log.Print("Create connection error:", err)
		return
	}

	defer func() {
		container.websocketConnectionFactory.GetInboundChannel().Disconnect(handler)
		if err != nil {
			log.Println("Error when close connection")
		}
	}()

	for {
		var req WSRequest
		err := conn.ReadJSON(&req)
		if err != nil {
			log.Println("Error when read json: ", err)
			return
		}

		//switch req.Action {
		//case SUBSCRIBE:
		//	client.Subscribe(req)
		//case UNSUBSCRIBE:
		//	client.Unsubscribe(req)
		//case BROADCAST:
		//	client.Broadcast(req)
		//case SEND_TO_USER:
		//	client.SendToUser(req)
		//default:
		//	fmt.Println("Action do not support")
		//}
	}
}
