package client

import (
	"errors"
	"net/http"
	"stress-testing/internal/biz"
)

func Request(sr *biz.StressRequest) (resp *http.Response, requestTime uint64, err error) {
	// todo
	err =errors.New("todo")
	return
}
