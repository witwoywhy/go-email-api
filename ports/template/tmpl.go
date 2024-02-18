package tmpl

var Template map[string]EventTemplate = map[string]EventTemplate{}

type EventTemplate struct {
	Subject string
	Body    string
}
