package tmpl

import (
	"fmt"
	"io"
	"os"
)

type GetTemplateFunc func(string) (*EventTemplate, error)

var GetTemplate GetTemplateFunc = func(event string) (*EventTemplate, error) {
	eventTemplate, ok := Template[event]
	if ok {
		return &eventTemplate, nil
	}

	subject, err := readFile(fmt.Sprintf("./template/%s/subject.tmpl", event))
	if err != nil {
		return nil, err
	}

	body, err := readFile(fmt.Sprintf("./template/%s/body.tmpl", event))
	if err != nil {
		return nil, err
	}

	eventTemplate = EventTemplate{
		Subject: string(subject),
		Body:    string(body),
	}
	Template[event] = eventTemplate

	return &eventTemplate, nil
}

func readFile(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return io.ReadAll(f)
}
