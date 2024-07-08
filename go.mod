module github.com/matt-primrose/gwm-tester

go 1.20

replace github.com/open-amt-cloud-toolkit/go-wsman-messages/v2 => ../go-wsman-messages

require github.com/open-amt-cloud-toolkit/go-wsman-messages/v2 v2.5.2

require (
	github.com/google/uuid v1.6.0 // indirect
	github.com/gorilla/websocket v1.5.2 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	golang.org/x/net v0.23.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
