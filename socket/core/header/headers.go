package header

const (
	StompIdHeader            = "id"
	StompHostHeader          = "host"
	StompAcceptVersionHeader = "accept-version"
	StompMessageIdHeader     = "stompmsg-id"
	StompReceiptHeader       = "receipt"
	StompReceiptIdHeader     = "receipt-id"
	StompSubscriptionHeader  = "subscription"
	StompVersionHeader       = "version"
	StompMessageHeader       = "stompmsg"
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
	headers map[string]string
}

func NewHeader() *Headers {
	return &Headers{headers: make(map[string]string, 0)}
}

func (h *Headers) SetHeader(key string, value string) {
	// TODO HGA WILL CHECK KEY VALUE IN RANGE
	h.headers[key] = value
}

func (h *Headers) GetHeader(key string) string {
	return h.headers[key]
}

func (h *Headers) GetHeaderProperties() map[string]string {
	return h.headers
}
