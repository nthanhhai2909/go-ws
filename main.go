package main

import (
	"mem-ws/core"
	"mem-ws/core/conf"
	"mem-ws/core/conf/broker"
	"mem-ws/core/conf/connection"
)

func main() {
	brokerRegistry := broker.MessageBrokerRegistry{
		ApplicationDestinationPrefix: "/ws",
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
	wsStarter.Start()
}
