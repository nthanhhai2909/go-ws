package main

import (
	"flag"
	"fmt"
	"log"
	"mem-ws/core"
	"mem-ws/core/conf"
	"mem-ws/core/conf/broker"
	"mem-ws/core/conf/connection"
	"net/http"
)

var addr = flag.String("addr", "localhost:8999", "http service address")

func main() {
	brokerRegistry := broker.MessageBrokerRegistry{
		ApplicationDestinationPrefixes: []string{"/ws"},
		StompBrokerRegistration: &broker.StompBrokerRegistration{
			ReplayHost:     "localhost",
			ReplayPort:     61613,
			ClientLogin:    "admin",
			ClientPasscode: "admin",
			Destinations:   []string{"/user/queue/1", "/user/queue/2", "/user/queue/2"},
		}}
	wsStarter := core.NewWSStarter(conf.Configuration{
		ConnectionConf: connection.NewDefaultConfiguration(),
		BrokerRegistry: brokerRegistry,
	})
	http.HandleFunc("/ws", wsStarter.Handler)
	fmt.Println("Server start listening at: localhost:8999")
	log.Fatal(http.ListenAndServe(*addr, nil))
}
