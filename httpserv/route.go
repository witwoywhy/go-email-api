package httpserv

import (
	"email-api/handler"
	"email-api/infrastructure"
	"email-api/ports/getfile"
	"email-api/ports/sendout"
	"email-api/services/send"

	"github.com/gin-gonic/gin"
)

func bindSendRoute(app *gin.Engine) {
	getFile := getfile.NewAdaptorMinio(infrastructure.MinioClient)
	sendOut := sendout.NewAdaptorMail(infrastructure.Mailer)

	service := send.New(getFile, sendOut)
	handle := handler.NewSendHandler(service)

	app.POST("/v1/notification", handle.Handle)
}
