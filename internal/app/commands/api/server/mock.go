package server

import "github.com/stretchr/testify/mock"

type Mock struct {
	mock.Mock
}

func (m *Mock) Start(address string) error {
	args := m.Called(address)
	return args.Error(0)
}
