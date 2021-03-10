package httpclient

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"

	"btc-pool-appserver/application/library/log"
)

var (
	MaxRetry    = 3
	MaxTimeOut  = 2000 //毫秒
	EmptyBuffer = &bytes.Buffer{}
)

func Get(url string, headers map[string]string, timeout int, retry int) ([]byte, error) {
	return requestWithTrys("GET", url, nil, headers, timeout, retry)
}

func Post(url string, params *bytes.Buffer, headers map[string]string, timeout int, retry int) ([]byte, error) {
	return requestWithTrys("POST", url, params, headers, timeout, retry)
}

func requestWithTrys(method, url string, params *bytes.Buffer, headers map[string]string, timeout int, retry int) ([]byte, error) {
	var result []byte
	var err error

	if retry < 0 {
		retry = MaxRetry
	}

	for i := 0; i < retry+1; i++ {
		result, err = Request(method, url, *params, headers, timeout)
		if err == nil {
			break
		}
		log.Errorf("requestWithTrys %v: %s", i, err.Error())
	}

	if err != nil {
		return nil, err
	}

	return result, nil
}

func Request(method, url string, params bytes.Buffer, headers map[string]string, timeout int) ([]byte, error) {

	log.Info(method, " api: ", url, " params:", params.String())

	var req *http.Request
	var err error

	if method == "GET" {
		req, err = http.NewRequest("GET", url, nil)
	} else {
		req, err = http.NewRequest("POST", url, &params)
	}
	if err != nil {
		return nil, err
	}

	if method == "POST" {
		req.Header.Set("Content-Type", "application/json")
	}

	// 设置header
	for k, v := range headers {
		if k == "Host" {
			req.Host = v
		} else {
			req.Header.Set(k, v)
		}
	}

	// 输出到 log
	//byts, _ := httputil.DumpRequest(req, true)
	//log.Info(string(byts))

	if timeout <= 0 {
		timeout = MaxTimeOut
	}
	c := &http.Client{
		Timeout: time.Millisecond * time.Duration(timeout),
	}

	var res *http.Response
	if res, err = c.Do(req); err != nil {
		return nil, err
	}
	defer func() {
		_ = res.Body.Close()
	}()

	var body []byte
	if body, err = ioutil.ReadAll(res.Body); err != nil {
		return nil, err
	}

	log.Info("post api res: ", string(body))
	return body, nil
}
