package main

import (
	"flag"
	"fmt"
	"log"
	"mem-ws/socket/core/conf"
	"mem-ws/socket/core/stomp"
	"net/http"
)

var addr = flag.String("addr", "localhost:8999", "http service address")

func main() {
	wsStarter := stomp.NewWSStarter(conf.NewDefaultWebsocketConnectionConfiguration())
	http.HandleFunc("/ws", wsStarter.Handler)
	fmt.Println("Server start listening at: localhost:8999")
	log.Fatal(http.ListenAndServe(*addr, nil))
}
