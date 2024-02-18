package sendout

import "github.com/stretchr/testify/mock"

type Mock struct {
	mock.Mock
}

func NewMock() *Mock {
	return &Mock{}
}

func (m *Mock) Execute(request Request) error {
	args := m.Called(request)
	return args.Error(0)
}
