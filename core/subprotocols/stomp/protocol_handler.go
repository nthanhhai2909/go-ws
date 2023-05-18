package stomp

import (
	"fmt"
	"github.com/nthanhhai2909/go-commons-lang/sliceutils"
	"github.com/nthanhhai2909/go-commons-lang/stringutils"
	"mem-ws/core/conf/broker"
	"mem-ws/core/errors"
	"mem-ws/core/subprotocols/stomp/channel/inbound"
	"mem-ws/core/subprotocols/stomp/cmd/client"
	"mem-ws/core/subprotocols/stomp/cmd/server"
	"mem-ws/core/subprotocols/stomp/codec"
	"mem-ws/core/subprotocols/stomp/constants"
	"mem-ws/core/subprotocols/stomp/header"
	"mem-ws/core/subprotocols/stomp/smsg"
	"mem-ws/native/message"
	"mem-ws/native/session"
	"mem-ws/native/subprotocol"
	"strings"
	"sync"
)

// ProtocolHandler - socket.IWebsocketHandler Implementation
type ProtocolHandler struct {
	Decoder        *codec.Decoder
	Encoder        *codec.Encoder
	InboundManager *inbound.Manager
}

func NewProtocolHandler(registration *broker.StompBrokerRegistration) subprotocol.ISubProtocolHandler {
	ibManager := &inbound.Manager{InboundMap: sync.Map{}}
	for _, destination := range registration.Destinations {
		ibManager.InboundMap.Store(destination, &inbound.Subscribable{Subscribers: &sync.Map{}})
	}
	return &ProtocolHandler{
		Decoder:        &codec.Decoder{},
		Encoder:        &codec.Encoder{},
		InboundManager: ibManager,
	}
}

func (h *ProtocolHandler) SupportProtocols() []string {
	return constants.SupportVersion
}

// HandleMessageFromClient TODO BROKER PROCESS
func (h *ProtocolHandler) HandleMessageFromClient(session session.ISession, message message.IMessage) {
	encoder := h.Encoder
	decoder := h.Decoder
	ibManager := h.InboundManager
	msg, _ := decoder.Decode(message.GetPayload())
	headers := msg.GetMessageHeaders()
	destination := headers.Destination()
	cmd := headers.Command()
	switch cmd {
	case client.Connect:
		stompVersion, err := commonVersionUse(headers.AcceptVersion())
		if err != nil {
			session.SendMessage(encoder.Encode(smsg.Error(map[string]string{
				header.StompVersion:     constants.SupportVersionInString,
				header.StompContentType: constants.TextPlain,
			}, []byte(err.Error()))))
			return
		}
		session.SendMessage(encoder.Encode(smsg.Connected(stompVersion)))
	case client.Send:
		if stringutils.IsBlank(destination) {
			session.SendMessage(encoder.Encode(smsg.Error(map[string]string{
				header.StompVersion:       constants.SupportVersionInString,
				header.StompDestination:   destination,
				header.StompContentLength: constants.TextPlain,
			}, []byte("Destination must not be null"))))
		}
		ibManager.Send(destination, msg)
	case client.Subscribe:
		err := ibManager.Subscribe(msg, session)
		if err != nil {
			// TODO HANDLER MESSAGE LATER
			session.SendMessage(encoder.Encode(smsg.Error(map[string]string{}, []byte(err.Error()))))
			return
		}
		session.SendMessage(encoder.Encode(smsg.NewGenericMessage(map[string]string{
			header.Command:           server.Message,
			header.StompDestination:  msg.GetMessageHeaders().Destination(),
			header.StompContentType:  constants.TextPlain,
			header.StompSubscription: msg.GetMessageHeaders().ID(),
		}, []byte("Welcome to VN"))))

	case client.Unsubscribe:
		err := ibManager.UnSubscribe(msg, session)
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

func (h *ProtocolHandler) HandleConnectionClose(session session.ISession) {
	err := h.InboundManager.HandleConnectionClose(session)
	if err != nil {
		// TODO HGA
	}
}
func commonVersionUse(clientAcceptVersion string) (string, error) {
	var clientVersions []string
	if stringutils.IsBlank(clientAcceptVersion) {
		clientVersions = []string{"v10.stomp"}
	} else {
		clientVersions = strings.Split(clientAcceptVersion, ",")
	}

	commons := sliceutils.Union(clientVersions, constants.SupportVersion)

	if sliceutils.IsEmpty(commons) {
		return stringutils.EMPTY, errors.MessageConversion{Message: fmt.Sprintf("Supported protocol versions are %s", constants.SupportVersionInString)}
	}
	return commons[0], nil
}
