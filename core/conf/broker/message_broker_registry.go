package broker

type MessageBrokerRegistry struct {
	ApplicationDestinationPrefixes []string
	StompBrokerRegistration        *StompBrokerRegistration
	InMemoryBrokerRegistration     *InMemoryBrokerRegistration
}
