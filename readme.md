# fakeSMTPserver

SMTP server for local development. Capture all incoming SMTP emails and send them to standard output. Quick and simple.

### Installation

#### For general users
```
curl -sf https://gobinaries.com/enriquebris/fakesmtpserver | sh
```

#### For golang users
Pre-requisite: golang

```
go get github.com/enriquebris/fakesmtpserver
cd $(go env GOPATH)/src/github.com/enriquebris/fakesmtpserver
go build
```

After *go build* the binary will be reachable from $(go env GOPATH)/bin

### Usage

Start listening on port 2025
```
fakesmtpserver start --addr :2525
```

Start listening on port 2025, no auth
```
fakesmtpserver start --addr :2525 --allowInsecureAuth true
```