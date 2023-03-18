package stomp

import (
	"fmt"
	"github.com/nthanhhai2909/go-commons-lang/sliceutils"
	"github.com/nthanhhai2909/go-commons-lang/stringutils"
	"mem-ws/core/conf/broker"
	"mem-ws/core/errors"
	"mem-ws/core/subprotocols/stomp/channel"
	"mem-ws/core/subprotocols/stomp/cmd/client"
	"mem-ws/core/subprotocols/stomp/constans"
	"mem-ws/core/subprotocols/stomp/smsg"
	"mem-ws/native/message"
	"mem-ws/native/session"
	"mem-ws/native/subprotocol"
	"strings"
)

// ProtocolHandler - socket.IWebsocketHandler Implementation
type ProtocolHandler struct {
	Decoder        *Decoder
	Encoder        *Encoder
	InboundManager *channel.InboundManager
}

func NewProtocolHandler(registration *broker.StompBrokerRegistration) subprotocol.ISubProtocolHandler {
	ibManager := &channel.InboundManager{InboundMap: make(map[string]channel.Inbound)}
	for _, destination := range registration.Destinations {
		ibManager.InboundMap[destination] = &channel.Subscribable{Subscribers: make(map[string]session.ISession)}
	}
	return &ProtocolHandler{
		Decoder:        &Decoder{},
		Encoder:        &Encoder{},
		InboundManager: ibManager,
	}
}

func (h *ProtocolHandler) SupportProtocols() []string {
	return constans.SupportVersion
}

// HandleMessageFromClient TODO BROKER PROCESS
func (h *ProtocolHandler) HandleMessageFromClient(session session.ISession, message message.IMessage) {
	encoder := h.Encoder
	decoder := h.Decoder
	ibManager := h.InboundManager
	msg, _ := decoder.Decode(message.GetPayload())
	headers := msg.GetMessageHeaders()
	destination := headers.GetHeader(constans.StompDestinationHeader)
	switch headers.GetHeader(constans.CommandHeader) {
	case client.Connect:
		stompVersion, err := commonVersionUse(headers.GetHeader(constans.StompAcceptVersionHeader))
		if err != nil {
			session.SendMessage(encoder.Encode(smsg.Error(map[string]string{
				constans.StompVersionHeader:     constans.SupportVersionInString,
				constans.StompContentTypeHeader: TextPlain,
			}, []byte(err.Error()))))
			return
		}
		session.SendMessage(encoder.Encode(smsg.Connected(stompVersion)))
		// TODO HGA
	case client.Send:
		fmt.Println("destination: ", destination)
		fmt.Println("msg: ", msg.GetPayload())
		if stringutils.IsBlank(destination) {
			session.SendMessage(encoder.Encode(smsg.Error(map[string]string{
				constans.StompVersionHeader:     constans.SupportVersionInString,
				constans.StompContentTypeHeader: TextPlain,
			}, []byte("Destination must not be null"))))
		}
	case client.Subscribe:
		fmt.Println("Subscribe: ", destination)
		fmt.Println("payload: ", msg.GetPayload())
		fmt.Println("Holder: ", ibManager.InboundMap)
		err := ibManager.Subscribe(destination, session)
		if err != nil {
			// TODO HANDLER MESSAGE LATER
			fmt.Println("hahahahah")
			session.SendMessage(encoder.Encode(smsg.Error(map[string]string{}, []byte(err.Error()))))
			return
		}
	case client.Unsubscribe:
		fmt.Println("Unsubscribe: ", destination)
		fmt.Println("payload: ", msg.GetPayload())
		fmt.Println("Holder: ", ibManager.InboundMap)
		err := ibManager.UnSubscribe(destination, session)
		if err != nil {
			// TODO HANDLER MESSAGE LATER
			session.SendMessage(encoder.Encode(smsg.Error(map[string]string{}, []byte(err.Error()))))
			return
		}
	}

}

// SendMessageToClient TODO IMPL - USE IN CASE SERVER ACTIVELY SEND MESSAGE TO CLIENT OR NOTIFICATION
func (h *ProtocolHandler) SendMessageToClient(session session.ISession, message smsg.IMessage) {
}

func commonVersionUse(clientAcceptVersion string) (string, error) {
	var clientVersions []string
	if stringutils.IsBlank(clientAcceptVersion) {
		clientVersions = []string{"v10.stomp"}
	} else {
		clientVersions = strings.Split(clientAcceptVersion, ",")
	}

	commons := sliceutils.Union(clientVersions, constans.SupportVersion)

	if sliceutils.IsEmpty(commons) {
		return stringutils.EMPTY, errors.MessageConversion{Message: fmt.Sprintf("Supported protocol versions are %s", constans.SupportVersionInString)}
	}
	return commons[0], nil
}
