package apiserver

import (
	"io"
	"net/http"

	"github.com/cvenkman/serv/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// APIserver struct ...
type APIserver struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store *store.Store
}

// New ...
func New(config *Config) *APIserver {
	return &APIserver{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start ...
func (s *APIserver) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}
	s.configureRouter()
	if err := s.configureStore(); err != nil {
		return err
	}

	s.logger.Info("start server")
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIserver) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)

	if err != nil {
		return err
	}

	s.logger.SetLevel(level)
	return nil
}

func (server *APIserver) configureRouter() {
	server.router.HandleFunc("/hello", server.handleHello())
}

func ( *APIserver) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hi")
	}
}

func (server *APIserver) configureStore() error {
	new_store := server.store.New(server.config.Store)
	if err := new_store.Open(); err != nil {
		return err
	}

	server.store = new_store
	return nil
}