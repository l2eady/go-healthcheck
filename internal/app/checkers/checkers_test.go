package checkers

import (
	"errors"
	"go-healthcheck/internal/app/lhttp/mocks"
	"testing"
)

func TestPingExpectSuccessReturnSuccess(t *testing.T) {
	url := "https://linecorp.com"
	expected := true
	mockCaller := &mocks.MockLHTTPCaller{MockGETReturnErr: nil}

	checker := checkerImpl{Caller: mockCaller}
	ok, err := checker.Ping(url)

	if err != nil {
		t.Errorf("Expected nil, but got %v", err)
	}
	if !ok {
		t.Errorf("Expected %v, but got %v", expected, ok)
	}
}

func TestPingExpectFailedReturnErrTimeOut(t *testing.T) {
	url := "https://linecorp.com"
	mockErr := errors.New("request timeout")
	expected := mockErr
	mockCaller := &mocks.MockLHTTPCaller{MockGETReturnErr: mockErr}

	checker := checkerImpl{Caller: mockCaller}
	ok, err := checker.Ping(url)

	if err == nil {
		t.Errorf("Expected %v, but got %v", expected, err)
	}
	if ok {
		t.Errorf("Expected false, but got %v", ok)
	}
}
