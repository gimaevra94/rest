package server

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Server struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

func New(config *Config) *Server {
	return &Server{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (start *Server) Start() error {
	if err := start.loggConfiger(); err != nil {
		return err
	}
	start.RoutConfiger()
	start.logger.Info("Startiing API server")
	return http.ListenAndServe(start.config.Addr, start.router)
}

func (server Server) loggConfiger() error {
	level, err := logrus.ParseLevel(server.config.logLevel)
	if err != nil {
		return err
	}
	server.logger.SetLevel(level)
	return nil
}

func (server *Server) RoutConfiger() {
	server.router.HandleFunc("/пошел нахуй", server.handleHello())
}

func (server *Server) handleHello() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer, "пошел нахуй")
	}
}
