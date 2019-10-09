package checkers

import (
	"fmt"
	"go-healthcheck/internal/app/lhttp"
	"net/http"
	"time"
)

type CheckerService interface {
	Ping(url string) (bool, error)
}
type checkerServiceImpl struct {
	Caller lhttp.HttpCaller
}

func New(maxTimeOut time.Duration) CheckerService {
	return &checkerServiceImpl{
		Caller: &lhttp.Caller{
			Body:   nil,
			Header: map[string]string{},
			Client: &http.Client{Timeout: maxTimeOut},
		},
	}
}

func (checker checkerServiceImpl) Ping(url string) (ok bool, err error) {
	checker.Caller.SetURL(url)
	if _, err = checker.Caller.GET(); err != nil {
		fmt.Printf("Ping to %s failed\n", url)
		return
	}
	ok = true
	fmt.Printf("Ping to %s successfuly\n", url)
	return
}
