package broker

type StompBrokerRegistration struct {
	Destinations   []string
	ReplayHost     string
	ReplayPort     int
	ClientLogin    string
	ClientPasscode string
}
