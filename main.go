package main

import (
	"flag"
	"log"
	"mem-ws/wscore"
	"mem-ws/wscore/handlers"
	"net/http"
)

var addr = flag.String("addr", "localhost:8999", "http service address")

func main() {
	upgrader := wscore.DefInstance()
	hub := handlers.NewHub()
	go hub.Start()
	wsService := handlers.New(upgrader, hub)
	http.HandleFunc("/ws", wsService.Handler)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
