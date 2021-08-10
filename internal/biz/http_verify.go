package biz

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func HttpStatusCode(request *StressRequest, response *http.Response) (responseCode int, isSuccessed bool) {

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

//
func RegisterVerifyHttp(verify string,verifyFunc VerifyHttp)  {

	defer  VerifyMapHttpMutex.Unlock()
	VerifyMapHttpMutex.Lock()
	key := fmt.Sprintf("%s.%s", FormTypeHTTP, verify)
	VerifyMapHttp[key] = verifyFunc
}
