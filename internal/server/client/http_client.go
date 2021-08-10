package client

import (
	"errors"
	"net/http"
	"stress-testing/internal/biz"
	"time"
)

func Request(sr *biz.StressRequest) (resp *http.Response, requestTime uint64, err error) {
	// todo
	time.Sleep(time.Second)
	err =errors.New("todo")
	return
}
