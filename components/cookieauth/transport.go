package cookieauth

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"
)

var cookieCache = sync.Map{}

// Transport transport
type Transport struct {
	conf *Config

	cookie *http.Cookie

	transport http.RoundTripper
	mutex     sync.Mutex
}

// NewTransport 构造认证使用的 Transport
func NewTransport(conf *Config, trs ...http.RoundTripper) *Transport {
	tr := http.DefaultTransport
	if len(trs) > 0 {
		tr = trs[0]
	}

	return &Transport{
		conf:      conf,
		transport: tr,
		mutex:     sync.Mutex{},
	}
}

// RoundTrip roundtrip req
func (tr *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	cacheKey := tr.conf.Username
	cookieIter, ok := cookieCache.Load(cacheKey)
	if ok {
		tr.cookie = cookieIter.(*http.Cookie)
	}

	if tr.expired() {
		err := tr.updateCookie()
		if err != nil {
			return nil, err
		}

		cookieCache.Store(cacheKey, tr.cookie)
	}

	req.Header.Set("Cookie", tr.cookie.String())
	return tr.transport.RoundTrip(req)
}

func (tr *Transport) expired() bool {
	if tr.cookie == nil {
		return true
	}

	// 提前 10 分钟刷新 cookie
	return tr.cookie.Expires.After(time.Now().Add(-10 * time.Minute))
}

func (tr *Transport) updateCookie() (err error) {
	tr.mutex.Lock()
	defer tr.mutex.Unlock()

	if !tr.expired() {
		return
	}

	body := new(strings.Builder)
	body.WriteString(`{"username":"`)
	body.WriteString(tr.conf.Username)
	body.WriteString(`","password":"`)
	body.WriteString(tr.conf.Password)
	body.WriteString(`"}`)

	u := tr.conf.Endpoint + "/api/gaea/signin"
	req, err := http.NewRequest(http.MethodPost, u, strings.NewReader(body.String()))
	if err != nil {
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	err = checkError(resp)
	if err != nil {
		return
	}

	cookie := getCookieFromResponse(resp)
	if cookie == nil {
		err = errors.New("get cookie `" + cookieName + "` failed")
		return
	}

	tr.cookie = cookie

	return
}

const (
	xreqID     = "X-Reqid"
	cookieName = "PORTAL_SESSION"
)

func getCookieFromResponse(response *http.Response) *http.Cookie {
	for _, cookie := range response.Cookies() {
		if cookieName == cookie.Name {
			return cookie
		}
	}

	return nil
}

func checkError(response *http.Response) (err error) {
	defer response.Body.Close()

	ret := loginResp{}
	reqID := response.Header.Get(xreqID)

	if response.StatusCode/100 != 2 {
		err = errors.New("request_id: " + reqID + ", status: " + response.Status + ", status_code: " + fmt.Sprint(response.StatusCode))
		return
	}

	err = json.NewDecoder(response.Body).Decode(&ret)
	if err != nil {
		err = errors.New("request_id: " + reqID + ", " + err.Error())
		return
	}

	if ret.Code != http.StatusOK {
		err = errors.New("request_id: " + reqID + ", " + ret.Message)
		return
	}

	return
}

type loginResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
