package main

import (
	"flag"
	"log"
	"mem-ws/wscore"
	"mem-ws/wscore/conf"
	"net/http"
)

var addr = flag.String("addr", "localhost:8999", "http service address")

func main() {
	defaultConfig := conf.NewDefaultWebsocketConnectionConfiguration()
	factory, _ := wscore.NewWebSocketConnectionFactory(defaultConfig)
	container := wscore.NewWSContainer(factory)
	container.StartServe()
	http.HandleFunc("/ws", container.Handler)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
