package send_test

import (
	"email-api/ports/getfile"
	"email-api/ports/sendout"
	tmpl "email-api/ports/template"
	"email-api/services/send"
	"email-api/utils/errs"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestExecute(t *testing.T) {
	type details struct {
		AccountName string `json:"accountName"`
		AccountNo   string `json:"accountNo"`
	}
	type given struct {
		request send.Request
	}
	type getFile struct {
		response *getfile.Response
		err      error
	}
	type sendOut struct {
		err error
	}
	type when struct {
		getFile getFile
		sendOut sendOut

		getTemplateFunc tmpl.GetTemplateFunc
	}
	type expect struct {
		response *send.Response
		err      errs.AppError
	}
	type testCase struct {
		name string

		given  *given
		when   *when
		expect *expect
	}

	d := details{
		AccountName: "Tester Test",
		AccountNo:   "01111011",
	}

	testCases := []testCase{
		{
			name: "All Green",
			given: &given{
				request: send.Request{
					UserID:  "0f5384b4-cfe2-4e3e-95d4-3ba6b6d639ec",
					Email:   "mail@mail.com",
					Event:   "E10001",
					Details: d,
					Files:   []string{"E-2023-01-01.pdf"},
				},
			},
			when: &when{
				getFile: getFile{response: &getfile.Response{}},
				sendOut: sendOut{err: nil},
				getTemplateFunc: func(string) (*tmpl.EventTemplate, error) {
					return &tmpl.EventTemplate{}, nil
				},
			},
			expect: &expect{
				response: &send.Response{},
			},
		},
		{
			name: "Failed when send email out",
			given: &given{
				request: send.Request{
					UserID:  "0f5384b4-cfe2-4e3e-95d4-3ba6b6d639ec",
					Email:   "mail@mail.com",
					Event:   "E10001",
					Details: d,
					Files:   []string{"E-2023-01-01.pdf"},
				},
			},
			when: &when{
				getFile: getFile{response: &getfile.Response{}},
				sendOut: sendOut{err: errors.New("send failed")},
				getTemplateFunc: func(string) (*tmpl.EventTemplate, error) {
					return &tmpl.EventTemplate{}, nil
				},
			},
			expect: &expect{
				err: errs.New(http.StatusInternalServerError, errs.T001, ""),
			},
		},
		{
			name: "Failed when get file error",
			given: &given{
				request: send.Request{
					UserID:  "0f5384b4-cfe2-4e3e-95d4-3ba6b6d639ec",
					Email:   "mail@mail.com",
					Event:   "E10001",
					Details: d,
					Files:   []string{"E-2023-01-01.pdf"},
				},
			},
			when: &when{
				getFile: getFile{err: errors.New("get error")},
				sendOut: sendOut{},
				getTemplateFunc: func(string) (*tmpl.EventTemplate, error) {
					return &tmpl.EventTemplate{}, nil
				},
			},
			expect: &expect{
				err: errs.New(http.StatusInternalServerError, errs.T001, ""),
			},
		},
		{
			name: "Failed when get tmpl error",
			given: &given{
				request: send.Request{
					UserID:  "0f5384b4-cfe2-4e3e-95d4-3ba6b6d639ec",
					Email:   "mail@mail.com",
					Event:   "E10001",
					Details: d,
					Files:   []string{"E-2023-01-01.pdf"},
				},
			},
			when: &when{
				getFile: getFile{},
				sendOut: sendOut{},
				getTemplateFunc: func(string) (*tmpl.EventTemplate, error) {
					return nil, errors.New("get error")
				},
			},
			expect: &expect{
				err: errs.New(http.StatusInternalServerError, errs.T001, ""),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tmpl.GetTemplate = tc.when.getTemplateFunc

			getFile := getfile.NewMock()
			getFile.On("Execute", mock.Anything).Return(tc.when.getFile.response, tc.when.getFile.err)

			sendOut := sendout.NewMock()
			sendOut.On("Execute", mock.Anything).Return(tc.when.sendOut.err)

			service := send.New(getFile, sendOut)

			response, err := service.Execute(tc.given.request)
			if tc.expect.err != nil {
				assert.Equal(t, tc.expect.err.GetCode(), err.GetCode())
			}
			assert.Equal(t, tc.expect.response, response)
		})
	}
}
