package store

type Store struct {
	config *Config
}

func (store *Store) New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (store *Store) Open() error {
	return nil
}

func (store *Store) Close() {

}