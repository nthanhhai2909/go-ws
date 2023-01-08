package main

import (
	"flag"
	"fmt"
	"log"
	provider2 "mem-ws/socket/core/provider"
	"net/http"
)

var addr = flag.String("addr", "localhost:8999", "http service address")

func main() {
	defaultConfig := provider2.NewDefaultWebsocketConnectionConfiguration()
	factory, _ := provider2.NewWebSocketConnectionFactory(defaultConfig)
	container := provider2.NewWSContainer(factory)
	http.HandleFunc("/ws", container.Handler)
	fmt.Println("Server start listening at: localhost:8999")
	log.Fatal(http.ListenAndServe(*addr, nil))
}
