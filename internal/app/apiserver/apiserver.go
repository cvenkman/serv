package apiserver

// APIserver struct ...
type APIserver struct {
	config *Config
}

// New ...
func New(config *Config) *APIserver {
	return &APIserver{
		config: config,
	}
}

// Start ...
func (s *APIserver) Start() error {
	return nil
}

