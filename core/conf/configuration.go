package conf

import (
	"mem-ws/core/conf/broker"
	"mem-ws/core/conf/connection"
)

type Configuration struct {
	ConnectionConf connection.Configuration
	BrokerRegistry broker.MessageBrokerRegistry
}
