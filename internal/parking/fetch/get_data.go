package fetch

import (
	"app/internal/parking/model"
	"crypto/tls"

	"github.com/mailru/easyjson"
	"github.com/valyala/fasthttp"
)

var (
	req    *fasthttp.Request
	resp   *fasthttp.Response
	client *fasthttp.Client
)

func init() {
	req = fasthttp.AcquireRequest()
	resp = fasthttp.AcquireResponse()
	client = &fasthttp.Client{TLSConfig: &tls.Config{InsecureSkipVerify: true}}
}

// send request to apiString with fasthttp
// parse response with easyjson
// return result
func GetInfo(apiString string) (*model.Response, error) {

	req.SetRequestURI("http://data.egov.kz/api/v4/park_taxi/data?apiKey=" + apiString)
	req.Header.SetMethod("GET")

	if err := client.Do(req, resp); err != nil {
		return nil, err
	}

	var result model.Response

	if err := easyjson.Unmarshal(resp.Body(), &result); err != nil {
		return nil, nil
	}

	return &result, nil
}
