package main

import (
	"flag"
	"log"
	"mem-ws/core"
	"mem-ws/core/conf"
	"net/http"
)

var addr = flag.String("addr", "localhost:8999", "http service address")

func main() {
	defaultConfig := conf.NewDefaultWebsocketConnectionConfiguration()
	factory, _ := core.NewWebSocketConnectionFactory(defaultConfig)
	container := core.NewWSContainer(factory)
	container.StartServe()
	http.HandleFunc("/ws", container.Handler)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
