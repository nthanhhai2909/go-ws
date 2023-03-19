package broker

type MessageBrokerRegistry struct {
	ApplicationDestinationPrefix string
	StompBrokerRegistration      *StompBrokerRegistration
	InMemoryBrokerRegistration   *InMemoryBrokerRegistration
}
