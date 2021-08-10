package verify

import (
	"io"
	"io/ioutil"
	"net/http"
	"stress-testing/internal/biz"
)

func HttpStatusCode(request *biz.StressRequest, response *http.Response) (responseCode int, isSuccessed bool) {

	defer func() {
		_ = response.Body.Close()
	}()
	responseCode = response.StatusCode
	// 报文返回的状态码与参数一致 说明ok
	if responseCode == request.Code {
		isSuccessed = true
	}
	io.Copy(ioutil.Discard, response.Body)
	return
}
