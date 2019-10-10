package healthycheck

import (
	"errors"
	"fmt"
	"go-healthcheck/internal/app"
	"go-healthcheck/internal/app/lhttp/mocks"
	"go-healthcheck/internal/app/models"
	"strings"
	"testing"
	"time"
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
	if !resp.IsSuccess {
		t.Errorf("Expected %v, but got %v", expected, resp.IsSuccess)
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
	if resp.IsSuccess {
		t.Errorf("Expected false, but got %v", resp.IsSuccess)
	}
}

func TestReadCSVExpectPingSuccessReturnReportSuccess(t *testing.T) {
	url := "https://linecorp.com"
	data := fmt.Sprintf(`%s,line company`, url)
	reader := strings.NewReader(data)
	mockConf := &app.Configs{}
	mockCaller := &mocks.MockLHTTPCaller{MockGETReturnErr: nil}
	checker := healthyCheckServiceImpl{Caller: mockCaller,
		Conf: mockConf}
	report := checker.HealthyCheckEndPointFromCSVFile(reader, 1)
	expectTotalData := 1
	expectTotalSuccess := 1

	if len(report.Data) != expectTotalData {
		t.Errorf("Expected %v, but got %v", expectTotalData, len(report.Data))
	}
	firstReport := report.Data[0]
	if !firstReport.IsSuccess {
		t.Errorf("Expected true, but got %v", firstReport.IsSuccess)
	}
	if firstReport.URL != url {
		t.Errorf("Expected %v, but got %v", url, firstReport.URL)
	}
	if report.TotalSuccess != expectTotalSuccess {
		t.Errorf("Expected %v, but got %v", expectTotalSuccess, report.TotalSuccess)
	}
}

func TestSendReportExpectSendFailedReturnFailed(t *testing.T) {
	mockConf := &app.Configs{}
	expectErr := errors.New("http request timeout")
	mockCaller := &mocks.MockLHTTPCaller{MockGETReturnErr: nil, MockPOSTReturnErr: expectErr}
	checker := healthyCheckServiceImpl{Caller: mockCaller,
		Conf: mockConf}

	err := checker.SendReport(&models.HealthyCheckReport{})
	if err == nil {
		t.Errorf("Expected %v, but got %v", expectErr, err)
	}

}

func TestReadCSVExpectPingFailedReturnReportFailed(t *testing.T) {
	url := "https://linecorp.com"
	data := fmt.Sprintf(`%s,line company`, url)
	reader := strings.NewReader(data)
	mockCaller := &mocks.MockLHTTPCaller{MockGETReturnErr: errors.New("request timeout")}
	mockConf := &app.Configs{}
	checker := healthyCheckServiceImpl{Caller: mockCaller, Conf: mockConf}

	report := checker.HealthyCheckEndPointFromCSVFile(reader, 1)
	expectTotalData := 1
	exppectTotalFail := 1
	if len(report.Data) != expectTotalData {
		t.Errorf("Expected %v, but got %v", expectTotalData, len(report.Data))
	}
	firstReport := report.Data[0]
	if firstReport.IsSuccess {
		t.Errorf("Expected true, but got %v", firstReport.IsSuccess)
	}
	if firstReport.URL != url {
		t.Errorf("Expected %v, but got %v", url, firstReport.URL)
	}
	if report.TotalFailure != exppectTotalFail {
		t.Errorf("Expected %v, but got %v", exppectTotalFail, report.TotalFailure)
	}
}

func TestReadCSVExpectReadErrorReturnReport(t *testing.T) {
	url := "https://linecorp.com"
	data := fmt.Sprintf(`%s,line "company`, url)
	reader := strings.NewReader(data)
	mockCaller := &mocks.MockLHTTPCaller{MockGETReturnErr: nil}
	mockConf := &app.Configs{}
	checker := healthyCheckServiceImpl{Caller: mockCaller, Conf: mockConf}
	report := checker.HealthyCheckEndPointFromCSVFile(reader, 1)
	expectTotalData := 0
	if len(report.Data) != expectTotalData {
		t.Errorf("Expected %v, but got %v", expectTotalData, len(report.Data))
	}

}

func TestNewHealthyCheckServiceExpectServiceReturnService(t *testing.T) {
	service := NewHealthyCheckService(time.Second, nil)
	_, ok := service.(HealthyCheckService)
	if !ok {
		t.Errorf("Expected true, but got %v", ok)
	}
}
