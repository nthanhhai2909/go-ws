package socket

type WebsocketHandler interface {
	AfterConnectionEstablished()
}
