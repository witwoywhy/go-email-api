package send

import (
	"email-api/ports/getfile"
	"email-api/ports/sendout"
	tmpl "email-api/ports/template"
	"email-api/utils/errs"
	"encoding/json"
	"net/http"

	"github.com/aymerick/raymond"
)

func (request *Request) ToMapping() (map[string]any, error) {
	b, err := json.Marshal(request.Details)
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, errs.T001, "")
	}

	var mapping map[string]any
	err = json.Unmarshal(b, &mapping)
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, errs.T001, "")
	}

	return mapping, nil
}

type Render struct {
	Subject string
	Body    string
}

func renderTemplate(mapping map[string]any, template *tmpl.EventTemplate) (*Render, error) {
	subject, err := raymond.Render(template.Subject, mapping)
	if err != nil {
		return nil, err
	}

	body, err := raymond.Render(template.Body, mapping)
	if err != nil {
		return nil, err
	}

	return &Render{
		Subject: subject,
		Body:    body,
	}, nil
}

func (request *Request) BuildSendOutRequest(render *Render, attachments []sendout.Attachment) sendout.Request {
	return sendout.Request{
		To:          request.Email,
		Subject:     render.Subject,
		Body:        render.Body,
		Attachments: attachments,
	}
}

func (request *Request) ToGetFileRequest(i int) getfile.Request {
	return getfile.Request{
		BucketName: request.UserID,
		FileName:   request.Files[i],
	}
}

func (s *service) getFiles(request Request) ([]sendout.Attachment, error) {
	attachments := make([]sendout.Attachment, len(request.Files))

	for i, file := range request.Files {
		object, err := s.getFile.Execute(request.ToGetFileRequest(i))
		if err != nil {
			return nil, err
		}

		attachments[i] = sendout.Attachment{
			FileName: file,
			Response: object,
		}
	}

	return attachments, nil
}
