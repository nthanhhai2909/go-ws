package smsg

type Handler interface {
	HandleMessage(msg IMessage) error
}
