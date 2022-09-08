package main

import (
	"flag"
	"log"
	"mem-ws/services/ws"
	"mem-ws/wscore"
	"net/http"
)

var addr = flag.String("addr", "localhost:9999", "http service address")

func main() {
	upgrader := wscore.DefInstance()
	hub := wscore.NewHub()
	go hub.Start()
	wsService := ws.New(upgrader, hub)
	http.HandleFunc("/ws", wsService.Handler)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
