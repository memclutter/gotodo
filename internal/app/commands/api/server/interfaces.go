package server

type Server interface {
	Start(address string) error
}
