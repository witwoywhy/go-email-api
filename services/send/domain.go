package send

import "email-api/utils/errs"

type Service interface {
	Execute(request Request) (*Response, errs.AppError)
}

type Request struct {
	UserID  string   `json:"userId"`
	Email   string   `json:"email"`
	Event   string   `json:"event"`
	Details any      `json:"details"`
	Files   []string `json:"files"`
}

type Response struct{}
