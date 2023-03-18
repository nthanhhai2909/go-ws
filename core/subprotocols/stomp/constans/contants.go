package constans

import "strings"

const EndLineStringRune = '\n'

const TerminalByte byte = 0

var SupportVersion = []string{"v10.stomp", "v11.stomp", "v12.stomp"}
var SupportVersionInString = strings.Join(SupportVersion, ",")

const (
	StompIdHeader            = "id"
	StompHostHeader          = "host"
	StompAcceptVersionHeader = "accept-version"
	StompMessageIdHeader     = "smsg-id"
	StompReceiptHeader       = "receipt"
	StompReceiptIdHeader     = "receipt-id"
	StompSubscriptionHeader  = "subscription"
	StompVersionHeader       = "version"
	StompMessageHeader       = "smsg"
	StompAckHeader           = "ack"
	StompNackHeader          = "nack"
	StompLoginHeader         = "login"
	StompPasscodeHeader      = "passcode"
	StompDestinationHeader   = "destination"
	StompContentTypeHeader   = "content-type"
	StompContentLengthHeader = "content-length"
	StompHeartbeatHeader     = "heart-beat"
	CommandHeader            = "stompCommand"
	CredentialsHeader        = "stompCredentials"
)
