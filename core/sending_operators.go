package core

import "mem-ws/core/simp"

type SendingOperations interface {
	Send(destination string, message simp.Message[interface{}]) error
}
