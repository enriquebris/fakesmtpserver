# fakeSMTPserver

Local SMTP server for local development. Capture all SMTP emails in local.

### Installation

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