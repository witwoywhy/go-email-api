package main

import (
	"email-api/httpserv"
	"email-api/infrastructure"
)

func init() {
	infrastructure.InitConfig()
}

func main() {
	infrastructure.InitAppConfig()
	infrastructure.InitMailSender()
	infrastructure.InitStorage()
	httpserv.Run()
}
