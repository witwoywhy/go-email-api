package sendout

import "email-api/ports/getfile"

type Port interface {
	Execute(request Request) error
}

type Request struct {
	To          string
	Subject     string
	Body        string
	Attachments []Attachment
}

type Attachment struct {
	FileName string
	*getfile.Response
}
