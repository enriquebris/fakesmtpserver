package main

import (
	"errors"
	"io"
	"io/ioutil"
	"log"

	gosmtp "github.com/emersion/go-smtp"
)

// SMTPBackend implements SMTP server methods.
type SMTPBackend struct{}

func (bkd *SMTPBackend) Login(state *gosmtp.ConnectionState, username, password string) (gosmtp.Session, error) {
	return &Session{}, nil
}

func (bkd *SMTPBackend) AnonymousLogin(state *gosmtp.ConnectionState) (gosmtp.Session, error) {
	return &Session{}, nil
}

func (bkd *SMTPBackend) NewSession(_ gosmtp.ConnectionState, _ string) (gosmtp.Session, error) {
	return &Session{}, nil
}

// A Session is returned after EHLO.
type Session struct{}

func (s *Session) AuthPlain(username, password string) error {
	log.Printf("[authplain]\n")
	if username != "username" || password != "password" {
		return errors.New("Invalid username or password")
	}
	return nil
}

func (s *Session) Mail(from string, opts gosmtp.MailOptions) error {
	//log.Println("Mail from:", from)
	log.Printf("[from] :: %v\n", from)
	return nil
}

func (s *Session) Rcpt(to string) error {
	//log.Println("Rcpt to:", to)
	log.Printf("[to] :: %v\n", to)
	return nil
}

func (s *Session) Data(r io.Reader) error {
	if b, err := ioutil.ReadAll(r); err != nil {
		return err
	} else {
		//log.Println("Data:", string(b))
		log.Printf("[data] :: %v\n", string(b))
	}
	return nil
}

func (s *Session) Reset() {
	log.Print("--end--\n")
}

func (s *Session) Logout() error {
	return nil
}
