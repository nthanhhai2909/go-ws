package header

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

type Headers struct {
	Properties map[string]string
}

func EmptyHeader() *Headers {
	return &Headers{Properties: make(map[string]string, 0)}
}

func WithProperties(properties map[string]string) *Headers {
	return &Headers{properties}
}

func (h *Headers) AddHeader(key string, value string) {
	// TODO HGA WILL CHECK KEY VALUE IN RANGE
	h.Properties[key] = value
}

func (h *Headers) GetHeader(key string) string {
	return h.Properties[key]
}
