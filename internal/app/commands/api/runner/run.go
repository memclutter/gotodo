package runner

func (s *runnerStruct) Run() error {
	return s.server.Start(":9000")
}
