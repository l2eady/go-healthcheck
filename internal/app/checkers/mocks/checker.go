package mocks

type MockCheckerService struct {
	MockPing    bool
	MockPingErr error
}

func (m *MockCheckerService) Ping(url string) (ok bool, err error) {
	return m.MockPing, m.MockPingErr
}
