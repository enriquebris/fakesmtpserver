package main

import (
	"log"
	"time"

	gosmtp "github.com/emersion/go-smtp"
)

type Server struct {
	config Config
}

func NewServer(cfg Config) *Server {
	srvr := &Server{}
	srvr.initialize(cfg)

	return srvr
}

func (st *Server) initialize(config Config) {
	st.config = config
}

func (st *Server) ListenAndServe() error {
	be := &SMTPBackend{}

	s := gosmtp.NewServer(be)

	s.Addr = st.config.Address
	s.Domain = st.config.Domain
	s.ReadTimeout = 10 * time.Second
	s.WriteTimeout = 10 * time.Second
	s.MaxMessageBytes = 1024 * 1024
	s.MaxRecipients = 50

	if st.config.AllowInsecureAuth {
		s.AllowInsecureAuth = true
	}

	log.Printf("Starting fake SMTP server at %v\n\n", s.Addr)
	return s.ListenAndServe()
}
