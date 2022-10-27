package stomp

import (
	"fmt"
	"github.com/gorilla/websocket"
	"mem-ws/socket/stomp/msg"
	"time"
)

const (
	writeWait = 10 * time.Second
)

type BrokerMessageHandler struct {
	Conn     *websocket.Conn
	Outbound chan []byte
}

func NewBrokerMessageHandler(conn *websocket.Conn) msg.Handler[[]byte] {
	broker := &BrokerMessageHandler{
		Conn:     conn,
		Outbound: make(chan []byte),
	}
	return broker
}

func (broker BrokerMessageHandler) HandleMessage(msg msg.Message[[]byte]) error {
	return nil
}

func (broker BrokerMessageHandler) GetConn() *websocket.Conn {
	return broker.Conn
}

func (broker BrokerMessageHandler) GetOutboundChannel() chan []byte {
	return broker.Outbound
}

func (broker *BrokerMessageHandler) outbound() {
	fmt.Println("Client Outbound InboundChannel started listen")

	defer func() {
		err := broker.Conn.Close()
		if err != nil {
			return
		}
	}()

	for {
		select {
		case payload, ok := <-broker.Outbound:
			err := broker.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err != nil {
				// TODO HGA WILL ADAPT LATER
			}

			if !ok {
				err = broker.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				if err != nil {
					// TODO HGA WILL ADAPT LATER
				}
				return
			}

			w, err := broker.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			_, err = w.Write(payload)
			if err != nil {
				// TODO HGA WILL ADAPT LATER
			}

			if err := w.Close(); err != nil {
				return
			}
		}
	}
}
