package header

const (
	StompIdHeader            = "id"
	StompHostHeader          = "host"
	StompAcceptVersionHeader = "accept-version"
	StompMessageIdHeader     = "msg-id"
	StompReceiptHeader       = "receipt"
	StompReceiptIdHeader     = "receipt-id"
	StompSubscriptionHeader  = "subscription"
	StompVersionHeader       = "version"
	StompMessageHeader       = "msg"
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
	headers map[string]interface{}
}

func NewHeader() *Headers {
	return &Headers{headers: make(map[string]interface{}, 0)}
}
func (h *Headers) SetDestination(destination interface{}) {
	h.headers[StompDestinationHeader] = destination
}

func (h *Headers) SetContentType(contentType interface{}) {
	h.headers[StompContentTypeHeader] = contentType
}

func (h *Headers) SetCommand(cmd interface{}) {
	h.headers[CommandHeader] = cmd
}

func (h *Headers) SetHeader(key string, value interface{}) {
	// TODO HGA WILL CHECK KEY VALUE IN RANGE
	h.headers[key] = value
}
