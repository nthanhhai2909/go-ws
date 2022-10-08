package message

const (
	StompIdHeader            = "id"
	StompHostHeader          = "host"
	StompAcceptVersionHeader = "accept-version"
	StompMessageIdHeader     = "message-id"
	StompReceiptHeader       = "receipt"
	StompReceiptIdHeader     = "receipt-id"
	StompSubscriptionHeader  = "subscription"
	StompVersionHeader       = "version"
	StompMessageHeader       = "message"
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

type Headers[V interface{}] struct {
	headers map[string]V
}

func (h *Headers[V]) SetDestination(destination V) {
	h.headers[StompDestinationHeader] = destination
}

func (h *Headers[V]) SetContentType(contentType V) {
	h.headers[StompContentTypeHeader] = contentType
}
