package socket

import (
	"net/http"
)

type WSStarter interface {
	Handler(w http.ResponseWriter, r *http.Request)
	AddEndPoint(endpoint string, handler MessageHandler)
}
