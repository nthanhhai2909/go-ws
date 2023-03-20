package header

const (
	StompId            = "id"
	StompHost          = "host"
	StompAcceptVersion = "accept-version"
	StompMessageId     = "smsg-id"
	StompReceipt       = "receipt"
	StompReceiptId     = "receipt-id"
	StompSubscription  = "subscription"
	StompVersion       = "version"
	StompMessage       = "smsg"
	StompAck           = "ack"
	StompNack          = "nack"
	StompLogin         = "login"
	StompPasscode      = "passcode"
	StompDestination   = "destination"
	StompContentType   = "content-type"
	StompContentLength = "content-length"
	StompHeartbeat     = "heart-beat"
	Command            = "stompCommand"
	Credentials        = "stompCredentials"
)

const (
	Client           = "client"
	AUTO             = "auto"
	ClientIndividual = "client-individual"
)

type Header struct {
	properties map[string]string
}

func EmptyHeader() *Header {
	return &Header{properties: make(map[string]string, 0)}
}

func WithProperties(properties map[string]string) *Header {
	return &Header{properties}
}

func (h *Header) AddHeader(key string, value string) {
	h.properties[key] = value
}

func (h *Header) GetHeader(key string) string {
	return h.properties[key]
}

func (h *Header) Properties() map[string]string {
	return h.properties
}

func (h *Header) ID() string {
	return h.properties[StompId]
}

func (h *Header) Command() string {
	return h.properties[Command]
}

func (h *Header) Ack() string {
	return h.properties[StompAck]
}

func (h *Header) Destination() string {
	return h.properties[StompDestination]
}

func (h *Header) AcceptVersion() string {
	return h.properties[StompAcceptVersion]
}
