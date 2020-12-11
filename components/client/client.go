package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	net_url "net/url"
	"strings"
	"time"

	"github.com/qiniu/stack-go/components/auth"
	"github.com/qiniu/stack-go/components/log"
)

const defaultTimestampFormat = "2006-01-02 15:04:05.999999"
const requestIDKey = "x-reqid"

// Client API 请求客户端
type Client struct {
	endpoint string

	log log.Logger
	tr  http.RoundTripper
}

// New API 请求客户端
func New(log log.Logger, endpoint string, credential *auth.Credential) *Client {
	return &Client{
		endpoint: endpoint,
		log:      log,
		tr: auth.NewTransport(credential, &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   600 * time.Second,
				KeepAlive: 600 * time.Second,
				DualStack: true,
			}).DialContext,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		}),
	}
}

// Call 请求
// 数据以 json 格式发送和接收
func (c *Client) Call(httpMethod string, url string, queries interface{}, data interface{}, ret interface{}) (err error) {

	start := time.Now()

	// 构造请求
	rawBody := []byte{}
	if data != nil && httpMethod != http.MethodGet {
		rawBody, err = json.Marshal(data)
		if err != nil {
			return
		}
	}

	req, err := http.NewRequest(httpMethod, c.endpoint+url, bytes.NewReader(rawBody))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.ContentLength = int64(len(rawBody))

	if queries != nil {
		switch queries.(type) {
		case net_url.Values:
			req.URL.RawQuery = queries.(net_url.Values).Encode()
		case map[string][]string:
			req.URL.RawQuery = net_url.Values(queries.(map[string][]string)).Encode()
		default:
			transQ, err := transObject2Queries(queries)
			if err != nil {
				return errors.New("invalid queries")
			}
			req.URL.RawQuery = transQ.Encode()
		}
	}

	// 读取结果
	resp, err := c.tr.RoundTrip(req)
	if err != nil {
		return err
	}

	defer func() {
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()

		end := time.Now()
		c.log.Debugf("[%v] [STACK_GO] %13v\t%s\n",
			end.Format(defaultTimestampFormat),
			end.Sub(start),
			resp.Request.URL.String(),
		)
	}()

	if resp.StatusCode/100 == 2 || resp.StatusCode/100 == 3 {
		if ret != nil && resp.ContentLength != 0 {
			err = json.NewDecoder(resp.Body).Decode(ret)
			if err != nil {
				return
			}
		}
		return nil
	}

	return responseError(resp)
}

func responseError(resp *http.Response) (err error) {
	err = &Error{
		Code:      resp.StatusCode,
		Name:      resp.Status,
		Message:   "UNKNOWN",
		RequestID: resp.Header.Get("requestIDKey"),
	}

	if resp.ContentLength != 0 {
		if ct := resp.Header.Get("Content-Type"); strings.TrimSpace(strings.SplitN(ct, ";", 2)[0]) == "application/json" {
			json.NewDecoder(resp.Body).Decode(err)
		}
	}

	return err
}

// transObject2Queries 将对象转成 url.Values
// 注意：仅转换一层结构，key 为 对象属性的 json tag 表示（或默认）
func transObject2Queries(obj interface{}) (q net_url.Values, err error) {
	bts, err := json.Marshal(obj)
	if err != nil {
		return q, err
	}

	var m map[string]interface{}
	err = json.Unmarshal(bts, &m)
	if err != nil {
		return q, err
	}

	q = make(map[string][]string)

	for k, v := range m {
		if v != nil {
			q[k] = []string{fmt.Sprintf("%v", v)}
		}
	}

	return
}
