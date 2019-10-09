package mocks

import "net/http"

type MockLHTTPCaller struct {
	URL                     string
	Body                    interface{}
	Header                  map[string]string
	Client                  *http.Client
	Response                interface{}
	MockGETReturnSuccess    *http.Response
	MockGETReturnErr        error
	MockPOSTReturnSuccess   *http.Response
	MockPOSTReturnErr       error
	MockPUTReturnSuccess    *http.Response
	MockPUTReturnErr        error
	MockPATCHReturnSuccess  *http.Response
	MockPATCHReturnErr      error
	MockDELETEReturnSuccess *http.Response
	MockDELETEReturnErr     error
}

// GET mocking GET method
func (m *MockLHTTPCaller) GET() (*http.Response, error) {
	return m.MockGETReturnSuccess, m.MockGETReturnErr
}

// POST mocking POST method
func (m *MockLHTTPCaller) POST() (*http.Response, error) {
	return m.MockPOSTReturnSuccess, m.MockPOSTReturnErr
}

// PUT mocking PUT method
func (m *MockLHTTPCaller) PUT() (*http.Response, error) {
	return m.MockPUTReturnSuccess, m.MockPUTReturnErr
}

// PATCH mocking PATCH method
func (m *MockLHTTPCaller) PATCH() (*http.Response, error) {
	return m.MockPATCHReturnSuccess, m.MockPATCHReturnErr
}

// DELETE mocking DELETE method
func (m *MockLHTTPCaller) DELETE() (*http.Response, error) {
	return m.MockDELETEReturnSuccess, m.MockDELETEReturnErr
}

// SetClient mocking http client
func (m *MockLHTTPCaller) SetClient(client *http.Client) {
	m.Client = client
}

// SetURL mocking request url
func (m *MockLHTTPCaller) SetURL(url string) {
	m.URL = url
}

// SetBody mocking request body
func (m *MockLHTTPCaller) SetBody(body interface{}) {
	m.Body = body
}

// SetHeader mocking request header
func (m *MockLHTTPCaller) SetHeader(header map[string]string) {
	m.Header = header
}
