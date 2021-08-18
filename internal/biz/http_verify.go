package biz

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

//通过http 状态码判断
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

// 先判断code 在判断是否是json
func HttpResponseJson(request *StressRequest, response *http.Response) (responseCode int, isSuccessed bool) {
	defer func() {
		_ = response.Body.Close()
	}()
	responseCode = response.StatusCode
	// 报文返回的状态码与参数一致 说明ok
	if responseCode == request.Code {
		// 验证json
		all, _ := ioutil.ReadAll(response.Body)
		valid := json.Valid(all)
		if valid {
			isSuccessed = true
		}
	}
	io.Copy(ioutil.Discard, response.Body)
	return
}

//注册校验方法
func RegisterVerifyHttp(verify string, verifyFunc VerifyHttp) {

	defer VerifyMapHttpMutex.Unlock()
	VerifyMapHttpMutex.Lock()
	key := fmt.Sprintf("%s.%s", FormTypeHTTP, verify)
	VerifyMapHttp[key] = verifyFunc
}
