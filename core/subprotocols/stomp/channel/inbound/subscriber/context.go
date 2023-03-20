package subscriber

import "mem-ws/native/session"

type Context struct {
	Ack     string
	Session session.ISession
}
