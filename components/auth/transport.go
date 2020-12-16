package auth

import (
	"encoding/base64"
	"net/http"
)

// Transport 认证请求 Transport
// 负责为请求添加认证签名
type Transport struct {
	credential *Credential
	transport  http.RoundTripper
}

// NewTransport 构造认证使用 Transport
func NewTransport(credential *Credential, transport http.RoundTripper) *Transport {
	if transport == nil {
		transport = http.DefaultTransport
	}

	return &Transport{
		credential: credential,
		transport:  transport,
	}
}

// RoundTrip 处理请求
// @impl net/http.RoundTripper#RoundTrip
func (t *Transport) RoundTrip(r *http.Request) (resp *http.Response, err error) {

	// base64 sign
	rawSign, err := SignRequest(t.credential.SecretKey, r)
	if err != nil {
		return nil, err
	}

	sign := base64.URLEncoding.EncodeToString(rawSign)

	req, _ := http.NewRequest(r.Method, r.URL.String(), r.Body)

	// copy headers
	req.Header = r.Header
	req.Header.Add("Authorization", "Qiniu "+t.credential.AccessKey+":"+string(sign))

	resp, err = t.transport.RoundTrip(req)

	return
}
