package handler

import (
	"email-api/services/send"
	"net/http"

	"github.com/gin-gonic/gin"
)

type sendHandler struct {
	service send.Service
}

func NewSendHandler(service send.Service) *sendHandler {
	return &sendHandler{
		service: service,
	}
}

func (h *sendHandler) Handle(ctx *gin.Context) {
	var request send.Request
	if err := ctx.BindJSON(&request); err != nil {
		ctx.Status(http.StatusBadRequest)
	}

	response, svcErr := h.service.Execute(request)
	if svcErr != nil {
		ctx.JSON(svcErr.GetHttpStatus(), svcErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
