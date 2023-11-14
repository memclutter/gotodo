package runner

import "github.com/stretchr/testify/mock"

type Mock struct {
	mock.Mock
}

func (m *Mock) Run() error {
	args := m.Called()
	return args.Error(0)
}
