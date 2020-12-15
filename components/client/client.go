package client

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/qiniu/stack-go/components/bytes"

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
func (c *Client) Call(r *Request, ret interface{}) (err error) {

	start := time.Now()

	req, err := http.NewRequest(r.Method, c.endpoint+r.URL, bytes.NewReader(r.Body))
	if err != nil {
		return err
	}

	for k, v := range r.Headers {
		req.Header.Set(k, v)
	}

	req.ContentLength = int64(len(r.Body))

	req.URL.RawQuery = r.Queries.Encode()
	fmt.Println(req.URL.String())

	// 读取结果
	resp, err := c.tr.RoundTrip(req)
	if err != nil {
		return err
	}

	defer func() {
		_, _ = io.Copy(ioutil.Discard, resp.Body)
		_ = resp.Body.Close()

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
		RequestID: resp.Header.Get(requestIDKey),
	}

	if resp.ContentLength != 0 {
		if ct := resp.Header.Get("Content-Type"); strings.TrimSpace(strings.SplitN(ct, ";", 2)[0]) == "application/json" {
			_ = json.NewDecoder(resp.Body).Decode(err)
		}
	}

	return err
}
