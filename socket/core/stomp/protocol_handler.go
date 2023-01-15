package stomp

import (
	"fmt"
	"github.com/nthanhhai2909/go-commons-lang/sliceutils"
	"github.com/nthanhhai2909/go-commons-lang/stringutils"
	"mem-ws/socket"
	"mem-ws/socket/core/errors"
	"mem-ws/socket/core/header"
	"mem-ws/socket/core/stomp/channel"
	"mem-ws/socket/core/stomp/cmd/client"
	"mem-ws/socket/core/stomp/smsg"
	"strings"
)

// ProtocolHandler - socket.ISubProtocolHandler Implementation
type ProtocolHandler struct {
	Decoder        *Decoder
	Encoder        *Encoder
	InboundChannel channel.Channel
	OutboundChanel channel.Channel
}

func NewProtocolHandler() socket.ISubProtocolHandler {
	return &ProtocolHandler{
		Decoder:        &Decoder{},
		Encoder:        &Encoder{},
		InboundChannel: &channel.Subscribable{},
		// TODO HGA WILL DO IT
		OutboundChanel: nil,
	}
}

func (h *ProtocolHandler) SupportProtocols() []string {
	return SupportVersion
}

// TODO BROKER PROCESS
func (h *ProtocolHandler) HandleMessageFromClient(session socket.IWebsocketSession, message socket.WebsocketMessage) {
	encoder := h.Encoder
	decoder := h.Decoder
	msg, _ := decoder.Decode(message.GetPayload())
	headers := msg.GetMessageHeaders()
	destination := headers.GetHeader(header.StompDestinationHeader)
	switch headers.GetHeader(header.CommandHeader) {
	case client.Connect:
		// TODO BROKER PROCESS
		stompVersion, err := commonVersionUse(headers.GetHeader(header.StompAcceptVersionHeader))
		if err != nil {
			session.SendMessage(encoder.Encode(smsg.Error(map[string]string{
				header.StompVersionHeader:     SupportVersionInString,
				header.StompContentTypeHeader: TextPlain,
			}, []byte(err.Error()))))
			return
		}
		session.SendMessage(encoder.Encode(smsg.Connected(stompVersion)))
	case client.Send:
		if stringutils.IsBlank(destination) {
			fmt.Println("test")
			// TODO HGA WILL RETURN ERROR TO CLIENT
		}
		//h.InboundChannel.Send(messageBuilder.)
	}
}

func (h *ProtocolHandler) SendMessageToClient(session socket.IWebsocketSession, message smsg.IMessage[[]byte]) {
}

func commonVersionUse(clientAcceptVersion string) (string, error) {
	var clientVersions []string
	if stringutils.IsBlank(clientAcceptVersion) {
		clientVersions = []string{"v10.stomp"}
	} else {
		clientVersions = strings.Split(clientAcceptVersion, ",")
	}

	commons := sliceutils.Union(clientVersions, SupportVersion)

	if sliceutils.IsEmpty(commons) {
		return stringutils.EMPTY, errors.MessageConversion{Message: fmt.Sprintf("Supported protocol versions are %s", SupportVersionInString)}
	}
	return commons[0], nil
}
