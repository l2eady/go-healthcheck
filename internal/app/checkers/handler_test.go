package checkers

import (
	"errors"
	"go-healthcheck/internal/app/checkers/mocks"
	"strings"
	"testing"
)

func TestReadCSVExpectFoundDataReturnData(t *testing.T) {
	mockCheckerService := mocks.MockCheckerService{MockPing: true}
	handler := NewCheckerHandler(&mockCheckerService)
	data := `https://linecorp.com,line company`
	reader := strings.NewReader(data)

	err := handler.CSV(reader)
	if err != nil {
		t.Errorf("Expected nil, but got %v", err)
	}
}

func TestReadCSVExpectErrReturnErr(t *testing.T) {
	mockCheckerService := mocks.MockCheckerService{MockPing: true}
	handler := NewCheckerHandler(&mockCheckerService)
	data := `https://linecorp.com,line "company`
	reader := strings.NewReader(data)
	expectErr := errors.New(`parse error on line 1, column 26: bare " in non-quoted-field`)
	err := handler.CSV(reader)
	if expectErr.Error() != err.Error() {
		t.Errorf("Expected %v, but got %v", expectErr, err)
	}
}
