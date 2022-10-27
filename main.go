package main

import (
	"flag"
	"fmt"
	"log"
	"mem-ws/socket"
	"mem-ws/socket/conf"
	"net/http"
)

var addr = flag.String("addr", "localhost:8999", "http service address")

func main() {
	defaultConfig := conf.NewDefaultWebsocketConnectionConfiguration()
	factory, _ := socket.NewWebSocketConnectionFactory(defaultConfig)
	container := socket.NewWSContainer(factory)
	http.HandleFunc("/ws", container.Handler)
	fmt.Println("Server start listening at: localhost:8999")
	log.Fatal(http.ListenAndServe(*addr, nil))
}
