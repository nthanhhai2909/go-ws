package subscriber

import (
	"mem-ws/native/session"
	"sync"
)

type Context struct {
	Ack          string
	Session      session.ISession
	Subscription *sync.Map
}
