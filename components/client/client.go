package client

import (
	"encoding/json"
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

	var (
		start = time.Now()
		l     = c.log
	)

	req, err := http.NewRequest(r.Method, c.endpoint+r.URL, bytes.NewReader(r.Body))
	if err != nil {
		l.Errorf("[%s] [STACK_GO] failed to new http request %s\n", start.Format(defaultTimestampFormat), err.Error())
		return err
	}

	for k, v := range r.Headers {
		req.Header.Set(k, v)
	}

	req.ContentLength = int64(len(r.Body))

	req.URL.RawQuery = r.Queries.Encode()

	l.Infof("[%s] [STACK_GO] [START] %s %s\n",
		start.Format(defaultTimestampFormat),
		r.Method,
		req.URL.String(),
	)

	// 读取结果
	resp, err := c.tr.RoundTrip(req)
	if err != nil {
		l.Errorf("[%s] [STACK_GO] failed to make http request %s\n", time.Now().Format(defaultTimestampFormat), err.Error())
		return err
	}

	defer func() {
		_, err2 := io.Copy(ioutil.Discard, resp.Body)
		if err2 != nil {
			l.Warnf("[%s] [STACK_GO] failed to readout response %s\n", time.Now().Format(defaultTimestampFormat), err.Error())
		}
		err2 = resp.Body.Close()
		if err2 != nil {
			l.Warnf("[%s] [STACK_GO] failed to close response body %s\n", time.Now().Format(defaultTimestampFormat), err.Error())
		}

		end := time.Now()
		l.Infof("[%s] [STACK_GO] [END] [%8v] %s %s\n",
			end.Format(defaultTimestampFormat),
			end.Sub(start),
			r.Method,
			req.URL.String(),
		)
	}()

	if resp.StatusCode/100 == 2 || resp.StatusCode/100 == 3 {
		if ret != nil && resp.ContentLength != 0 {
			err = json.NewDecoder(resp.Body).Decode(ret)
			if err != nil {
				l.Errorf("[%s] [STACK_GO] failed to decode response %s, request_id: %s\n", time.Now().Format(defaultTimestampFormat), err.Error(), resp.Header.Get(requestIDKey))
				return err
			}
		}
		return nil
	}

	return responseError(l, resp)
}

func responseError(l log.Logger, resp *http.Response) (err error) {
	err = &Error{
		Code:      resp.StatusCode,
		Name:      resp.Status,
		Message:   "",
		RequestID: resp.Header.Get(requestIDKey),
	}

	if resp.ContentLength != 0 {

		decodeErr := &struct {
			RequestID string `json:"request_id"`
			Error     Error  `json:"error"`
		}{}

		if ct := resp.Header.Get("Content-Type"); strings.TrimSpace(strings.SplitN(ct, ";", 2)[0]) == "application/json" {
			bts, err2 := ioutil.ReadAll(resp.Body)
			if err2 != nil {
				l.Errorf("[%s] [STACK_GO] failed to read error response %s, request_id: %s\n", time.Now().Format(defaultTimestampFormat), err.Error(), resp.Header.Get(requestIDKey))
				return err
			}

			err2 = json.Unmarshal(bts, decodeErr)
			if err2 != nil {
				l.Errorf("[%s] [STACK_GO] failed to unmarshal error response %s, request_id: %s, raw data: %s\n", time.Now().Format(defaultTimestampFormat), err.Error(), resp.Header.Get(requestIDKey), string(bts))
				return err
			}

			return &Error{
				Code:      decodeErr.Error.Code,
				Name:      decodeErr.Error.Name,
				Message:   decodeErr.Error.Message,
				RequestID: decodeErr.RequestID,
			}
		}
	}

	return err
}
