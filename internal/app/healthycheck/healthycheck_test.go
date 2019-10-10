package healthycheck

import (
	"errors"
	"go-healthcheck/internal/app/lhttp/mocks"
	"go-healthcheck/internal/app/models"
	"testing"
)

func TestPingExpectSuccessReturnSuccess(t *testing.T) {
	req := models.HealthyCheckRequest{URL: "https://linecorp.com"}
	expected := true
	mockCaller := &mocks.MockLHTTPCaller{MockGETReturnErr: nil}

	checker := healthyCheckServiceImpl{Caller: mockCaller}
	resp, err := checker.Ping(req)

	if err != nil {
		t.Errorf("Expected nil, but got %v", err)
	}
	if !resp.Status {
		t.Errorf("Expected %v, but got %v", expected, resp.Status)
	}
}

func TestPingExpectFailedReturnErrTimeOut(t *testing.T) {
	req := models.HealthyCheckRequest{URL: "https://linecorp.com"}
	mockErr := errors.New("request timeout")
	expected := mockErr
	mockCaller := &mocks.MockLHTTPCaller{MockGETReturnErr: mockErr}

	checker := healthyCheckServiceImpl{Caller: mockCaller}
	resp, err := checker.Ping(req)

	if err == nil {
		t.Errorf("Expected %v, but got %v", expected, err)
	}
	if resp.Status {
		t.Errorf("Expected false, but got %v", resp.Status)
	}
}

func TestReadCSVExpectFoundDataReturnData(t *testing.T) {
	// mockCheckerService := mocks.MockHealthyCheckService{MockPing: true}
	// handler := NewCheckerHandler(&mockCheckerService)
	// data := `https://linecorp.com,line company`
	// reader := strings.NewReader(data)

	// err := handler.HealthyCheckFromCSVFile(reader)
	// if err != nil {
	// 	t.Errorf("Expected nil, but got %v", err)
	// }
}

func TestReadCSVExpectErrReturnErr(t *testing.T) {
	// mockCheckerService := mocks.MockHealthyCheckService{MockPing: true}
	// handler := NewCheckerHandler(&mockCheckerService)
	// data := `https://linecorp.com,line "company`
	// reader := strings.NewReader(data)
	// expectErr := errors.New(`parse error on line 1, column 26: bare " in non-quoted-field`)
	// err := handler.HealthyCheckFromCSVFile(reader)
	// if expectErr.Error() != err.Error() {
	// 	t.Errorf("Expected %v, but got %v", expectErr, err)
	// }
}
