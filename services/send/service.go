package send

import (
	"email-api/ports/getfile"
	"email-api/ports/sendout"
	tmpl "email-api/ports/template"
	"email-api/utils/errs"
	"log"
	"net/http"
)

type service struct {
	getFile getfile.Port
	sendOut sendout.Port
}

func New(
	getFile getfile.Port,
	sendOut sendout.Port,
) Service {
	return &service{
		getFile: getFile,
		sendOut: sendOut,
	}
}

func (s *service) Execute(request Request) (*Response, errs.AppError) {
	mapping, err := request.ToMapping()
	if err != nil {
		log.Printf("Failed to make mapping: %v", err)
		return nil, errs.New(http.StatusInternalServerError, errs.T001, "")
	}

	template, err := tmpl.GetTemplate(request.Event)
	if err != nil {
		log.Printf("Failed to get template: %v", err)
		return nil, errs.New(http.StatusInternalServerError, errs.T001, "")
	}

	render, err := renderTemplate(mapping, template)
	if err != nil {
		log.Printf("Failed to render template: %v", err)
		return nil, errs.New(http.StatusInternalServerError, errs.T001, "")
	}

	attachments, err := s.getFiles(request)
	if err != nil {
		log.Printf("Failed to render template: %v", err)
		return nil, errs.New(http.StatusInternalServerError, errs.T001, "")
	}

	err = s.sendOut.Execute(request.BuildSendOutRequest(render, attachments))
	if err != nil {
		log.Printf("Failed to send email: %v", err)
		return nil, errs.New(http.StatusInternalServerError, errs.T001, "")
	}

	return &Response{}, nil
}
