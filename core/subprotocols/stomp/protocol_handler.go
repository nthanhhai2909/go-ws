package stomp

import (
	"fmt"
	"github.com/nthanhhai2909/go-commons-lang/sliceutils"
	"github.com/nthanhhai2909/go-commons-lang/stringutils"
	"log"
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
	InboundChannel channel.Channel
	OutboundChanel channel.Channel
}

func NewProtocolHandler() subprotocol.ISubProtocolHandler {
	return &ProtocolHandler{
		Decoder:        &Decoder{},
		Encoder:        &Encoder{},
		InboundChannel: &channel.Subscribable{},
		// TODO HGA WILL DO IT
		OutboundChanel: nil,
	}
}

func (h *ProtocolHandler) SupportProtocols() []string {
	return constans.SupportVersion
}

// TODO BROKER PROCESS
func (h *ProtocolHandler) HandleMessageFromClient(session session.ISession, message message.IMessage) {
	encoder := h.Encoder
	decoder := h.Decoder
	msg, _ := decoder.Decode(message.GetPayload())
	headers := msg.GetMessageHeaders()
	destination := headers.GetHeader(constans.StompDestinationHeader)
	switch headers.GetHeader(constans.CommandHeader) {
	case client.Connect:
		// TODO BROKER PROCESS
		stompVersion, err := commonVersionUse(headers.GetHeader(constans.StompAcceptVersionHeader))
		if err != nil {
			session.SendMessage(encoder.Encode(smsg.Error(map[string]string{
				constans.StompVersionHeader:     constans.SupportVersionInString,
				constans.StompContentTypeHeader: TextPlain,
			}, []byte(err.Error()))))
			return
		}
		session.SendMessage(encoder.Encode(smsg.Connected(stompVersion)))
	case client.Send:
		if stringutils.IsBlank(destination) {
			session.SendMessage(encoder.Encode(smsg.Error(map[string]string{
				constans.StompVersionHeader:     constans.SupportVersionInString,
				constans.StompContentTypeHeader: TextPlain,
			}, []byte("Destination must not be null"))))
		}
		err := h.InboundChannel.Send(msg)
		if err != nil {
			// TODO HGA WILL PROCESS IN CASE FAILED
			log.Fatal("Fail to send message to client")
		}
	}
}

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
