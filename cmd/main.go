package main

import (
	"fmt"
	"log"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

func main() {
	t := Tester{}
	t.start()
}

type Tester struct {
	gwm       wsman.Messages
	classList []gwmClass
	methodMap map[string]func()
}

type gwmClass struct {
	className  string
	methodName string
}

func (t Tester) getGeneralSettings() {
	result, err := t.gwm.AMT.GeneralSettings.Get()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	log.Printf("GeneralSettings: %v\n", result.XMLOutput)
}

func (t Tester) getAuditLogRecords() {
	result, err := t.gwm.AMT.AuditLog.ReadRecords(700)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	log.Printf("AuditLogRecords: %v\n", result.JSON())
}

func (t Tester) start() {
	t.gwm = t.setupGWM()

	t.methodMap = map[string]func(){
		"AMT_GeneralSettings.Get":  t.getGeneralSettings,
		"AMT_AuditLog.ReadRecords": t.getAuditLogRecords,
	}

	classes := []gwmClass{
		// {className: "AMT_GeneralSettings", methodName: "Get"},
		{className: "AMT_AuditLog", methodName: "ReadRecords"},
	}

	for _, class := range classes {
		methodKey := fmt.Sprintf("%s.%s", class.className, class.methodName)

		if method, ok := t.methodMap[methodKey]; ok {
			method()
		} else {
			fmt.Printf("Method %s not found for class %s\n", class.methodName, class.className)
		}
	}
}

func (t Tester) setupGWM() wsman.Messages {
	clientParameters := client.Parameters{
		Target:            "",
		Username:          "",
		Password:          "",
		UseDigest:         true,
		UseTLS:            false,
		SelfSignedAllowed: false,
		LogAMTMessages:    true,
		IsRedirection:     false,
	}
	gwm := wsman.NewMessages(clientParameters)
	return gwm
}
