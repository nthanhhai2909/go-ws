package wserror

const (
	InvalidConfigurationMessage = "Invalid Configuration"
)

type ConfigurationError struct {
	message string
}

func (e ConfigurationError) Error() string { return e.message }

func InvalidConfigurationError() ConfigurationError {
	return ConfigurationError{message: InvalidConfigurationMessage}
}
